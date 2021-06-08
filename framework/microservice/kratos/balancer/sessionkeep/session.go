package sessionkeep

import (
	"context"
)

// Session 会话信息
type Session struct {
	ID    string
	Value string
}

// SessionMgr 会话管理
type SessionMgr interface {
	GetSessionIDFromContext(context.Context) (string, bool)
	GetSessionInfoFromStorage(context.Context, string) (*Session, error)
	SetSessionInfoToStorage(context.Context, string, *Session) error
	DelSessionInfoFromStorage(context.Context, string) error
	NewLocker(sessionID string) Locker
}

// Locker 分布式锁
type Locker interface {
	LockContext(context.Context) error
	UnlockContext(context.Context) error
}
