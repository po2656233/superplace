package cherryActor

import (
	cerror "superplace/logger/error"
)

var (
	ErrForbiddenToCallSelf       = cerror.Errorf("SendActorID cannot be equal to TargetActorID")
	ErrForbiddenCreateChildActor = cerror.Errorf("Forbidden create child actor")
	ErrActorIDIsNil              = cerror.Error("actorID is nil.")
)

const (
	LocalName  = "local"
	RemoteName = "remote"
)
