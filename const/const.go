package superConst

import (
	"fmt"
	cerror "github.com/po2656233/superplace/logger/error"
)

const (
	DOT              = "." //ActorPath的分隔符
	version          = "2.0.0"
	RequestFuncName  = "request"
	ResponseFuncName = "response"
	LocalName        = "local"
	RemoteName       = "remote"
)

var (
	ErrForbiddenToCallSelf       = cerror.Errorf("SendActorID cannot be equal to TargetActorID")
	ErrForbiddenCreateChildActor = cerror.Errorf("Forbidden create child actor")
	ErrActorIDIsNil              = cerror.Error("actorID is nil.")
)
var logo = `game sever framework @v%s`

func GetLOGO() string {
	return fmt.Sprintf(logo, Version())
}

func Version() string {
	return version
}
