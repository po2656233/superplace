package actor

import (
	cerror "github.com/po2656233/superplace/logger/error"
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
