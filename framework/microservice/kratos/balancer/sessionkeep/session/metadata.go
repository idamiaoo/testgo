package session

import (
	"google.golang.org/grpc/metadata"
)

const sessionIdKey = "balancer_session_key_id"

func MDWithSessionID(md metadata.MD, sessionID string) metadata.MD {
	md.Set(sessionIdKey, sessionID)
	return md
}

// SessionIDFromMD .
func SessionIDFromMD(md metadata.MD) (string, bool) {
	values := md.Get(sessionIdKey)
	if len(values) == 0 {
		return "", false
	}
	return values[0], true
}
