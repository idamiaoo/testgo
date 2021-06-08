package sessionkeep

import (
	"math/rand"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/resolver"
)

const Name = "session_keep"

func NewBuilder(mgr SessionMgr) balancer.Builder {
	return base.NewBalancerBuilder(Name, &keepPickerBuilder{sessionMgr: mgr}, base.Config{HealthCheck: true})
}

type keepPickerBuilder struct {
	sessionMgr SessionMgr
}

// Build 实现 base.PickerBuilder 的接口方法
func (b *keepPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}
	p := &keepPicker{
		sessionMgr: b.sessionMgr,
	}
	for sc, sci := range info.ReadySCs {
		p.subConns = append(p.subConns, &subConn{conn: sc, addr: sci.Address})
	}
	p.next = rand.Intn(len(p.subConns))
	return p
}

// subConn 连接信息
type subConn struct {
	conn balancer.SubConn
	addr resolver.Address
}

// keepPicker 会话保持负载均衡选择器
type keepPicker struct {
	sessionMgr SessionMgr
	subConns   []*subConn
	mu         sync.Mutex
	next       int
}

// Pick 实现 balancer.Picker 接口方法
func (p *keepPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	// 从 context 获取 sessionID
	sessionID, _ := p.sessionMgr.GetSessionIDFromContext(info.Ctx)
	// 没有会话ID, 直接使用轮询负载均衡
	if sessionID == "" {
		p.mu.Lock()
		sc := p.subConns[p.next]
		p.next = (p.next + 1) % len(p.subConns)
		p.mu.Unlock()
		return balancer.PickResult{SubConn: sc.conn}, nil
	}

	// 获取分布式锁
	mu := p.sessionMgr.NewLocker(sessionID)
	if err := mu.LockContext(info.Ctx); err != nil {
		return balancer.PickResult{}, err
	}

	defer func() {
		if err := mu.UnlockContext(info.Ctx); err != nil {

		}
	}()

	// 从缓存中获取会话信息
	sessionInfo, err := p.sessionMgr.GetSessionInfoFromStorage(info.Ctx, sessionID)
	if err != nil {
		return balancer.PickResult{}, err
	}

	// 缓存中有会话信息
	if sessionInfo != nil {
		// 地址列表中能找到,有效,直接返回
		for _, v := range p.subConns {
			if v.addr.Addr == sessionInfo.Value {
				return balancer.PickResult{SubConn: v.conn}, nil
			}
		}
		// 地址列表中找不到(原因服务重启了等等),删除会话缓存
		if err := p.sessionMgr.DelSessionInfoFromStorage(info.Ctx, sessionID); err != nil {
			return balancer.PickResult{}, err
		}
	}

	// 缓存中没有会话或者缓存中会话已失效,重新选择一个链接并添加会话缓存

	// 轮询选择
	p.mu.Lock()
	sc := p.subConns[p.next]
	p.next = (p.next + 1) % len(p.subConns)
	p.mu.Unlock()

	// 写入缓存
	sess := &Session{
		ID:    sessionID,
		Value: sc.addr.Addr,
	}
	if err := p.sessionMgr.SetSessionInfoToStorage(info.Ctx, sessionID, sess); err != nil {
		return balancer.PickResult{}, err
	}

	return balancer.PickResult{SubConn: sc.conn}, nil
}
