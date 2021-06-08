package sessionkeep

import (
	"context"
)

type Session struct {
	ID    string
	Value string
}

type SessionMgr interface {
	GetSessionIDFromContext(context.Context) (string, bool)
	GetSessionInfoFromStorage(context.Context, string) (*Session, error)
	SetSessionInfoToStorage(context.Context, string, *Session) error
	DelSessionInfoFromStorage(context.Context, string) error
	NewLocker(sessionID string) Locker
}

type Locker interface {
	LockContext(context.Context) error
	UnlockContext(context.Context) error
}
