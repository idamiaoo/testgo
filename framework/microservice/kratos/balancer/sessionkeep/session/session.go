package session

import (
	"context"
	"encoding/json"

	"github.com/pescaria/testgo/framework/microservice/kratos/balancer/sessionkeep"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"google.golang.org/grpc/metadata"
)

const keyBalancerSessionCachePrefix = "balancer:session:"

type LockerImpl struct {
	mutex *redsync.Mutex
}

func (mu *LockerImpl) LockContext(ctx context.Context) error {
	return mu.mutex.LockContext(ctx)
}

func (mu *LockerImpl) UnlockContext(ctx context.Context) error {
	_, err := mu.mutex.UnlockContext(ctx)
	return err
}

type SessionMgrImpl struct {
	rds redis.UniversalClient
	rs  *redsync.Redsync
}

func NewSessionMgr(rds redis.UniversalClient) sessionkeep.SessionMgr {
	pool := goredis.NewPool(rds)
	rs := redsync.New(pool)
	return &SessionMgrImpl{
		rds: rds,
		rs:  rs,
	}
}

func (s *SessionMgrImpl) GetSessionIDFromContext(ctx context.Context) (string, bool) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return "", false
	}
	return SessionIDFromMD(md)
}

func (s *SessionMgrImpl) GetSessionInfoFromStorage(ctx context.Context, sessionID string) (*sessionkeep.Session, error) {
	val, err := s.rds.Get(ctx, keyBalancerSessionCachePrefix+sessionID).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var session sessionkeep.Session
	if err := json.Unmarshal([]byte(val), &session); err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *SessionMgrImpl) DelSessionInfoFromStorage(ctx context.Context, sessionID string) error {
	return s.rds.Del(ctx, keyBalancerSessionCachePrefix+sessionID).Err()
}

func (s *SessionMgrImpl) SetSessionInfoToStorage(ctx context.Context, sessionID string, session *sessionkeep.Session) error {
	val, err := json.Marshal(session)
	if err != nil {
		return err
	}
	return s.rds.Set(ctx, keyBalancerSessionCachePrefix+sessionID, val, 0).Err()
}

func (s *SessionMgrImpl) NewLocker(sessionID string) sessionkeep.Locker {
	name := "session_keep_sync" + "_" + sessionID
	mu := s.rs.NewMutex(name)
	return &LockerImpl{mutex: mu}
}
