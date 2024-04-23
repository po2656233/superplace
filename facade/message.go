package facade

import (
	cerr "github.com/po2656233/superplace/logger/error"
	"strings"
	"sync"
	"time"

	cconst "github.com/po2656233/superplace/const"
	cstring "github.com/po2656233/superplace/extend/string"
	cproto "github.com/po2656233/superplace/net/proto"
)

type (
	Message struct {
		BuildTime    int64            // message build time(ms)
		PostTime     int64            // post to actor time(ms)
		Source       string           // 来源actor path
		Target       string           // 目标actor path
		targetPath   *ActorPath       // 目标actor path对象
		FuncName     string           // 请求调用的函数名
		Session      *cproto.Session  // session of gateway
		Args         interface{}      // 请求的参数
		Err          error            // 返回的错误
		ClusterReply IRespond         // 返回消息的接口
		IsCluster    bool             // 是否为集群消息
		ChanResult   chan interface{} //
	}

	IRespond interface {
		Respond(data []byte) error
	}

	// ActorPath = NodeID . ActorID
	// ActorPath = NodeID . ActorID . ChildID
	ActorPath struct {
		NodeID  string
		ActorID string
		ChildID string
	}
)

var (
	messagePool = &sync.Pool{
		New: func() interface{} {
			return new(Message)
		},
	}
)

func GetMessage() *Message {
	msg := messagePool.Get().(*Message)
	msg.BuildTime = time.Now().UnixMilli()
	return msg
}

func BuildMessage(source, target, funcName string, arg interface{}) *Message {
	message := GetMessage()
	message.Source = source
	message.Target = target
	message.FuncName = funcName
	message.Args = arg
	message.ChanResult = make(chan interface{})

	return message
}

func (p *Message) Recycle() {
	p.Source = ""
	p.Target = ""
	p.targetPath = nil
	p.FuncName = ""
	p.Args = nil
	p.Err = nil
	p.ClusterReply = nil
	p.IsCluster = false
	p.ChanResult = nil
	p.BuildTime = 0
	p.PostTime = 0
	messagePool.Put(p)
}

func (p *Message) TargetPath() *ActorPath {
	if p.targetPath == nil {
		p.targetPath, _ = ToActorPath(p.Target)
	}
	return p.targetPath
}

func (p *Message) IsReply() bool {
	return p.ClusterReply != nil
}

func (p *ActorPath) IsChild() bool {
	return p.ChildID != ""
}

func (p *ActorPath) IsParent() bool {
	return p.ChildID == ""
}

// String
func (p *ActorPath) String() string {
	return NewChildPath(p.NodeID, p.ActorID, p.ChildID)
}

func NewActorPath(nodeID, actorID, childID string) *ActorPath {
	return &ActorPath{
		NodeID:  nodeID,
		ActorID: actorID,
		ChildID: childID,
	}
}

func NewChildPath(nodeID, actorID, childID interface{}) string {
	if childID == "" {
		return NewPath(nodeID, actorID)
	}
	return cstring.ToString(nodeID) + cconst.DOT + cstring.ToString(actorID) + cconst.DOT + cstring.ToString(childID)
}

func NewPath(nodeID, actorID interface{}) string {
	return cstring.ToString(nodeID) + cconst.DOT + cstring.ToString(actorID)
}

func ToActorPath(path string) (*ActorPath, error) {
	if path == "" {
		return nil, cerr.ActorPathError
	}

	p := strings.Split(path, cconst.DOT)
	pLen := len(p)

	if pLen == 2 {
		return NewActorPath(p[0], p[1], ""), nil
	}

	if len(p) == 3 {
		return NewActorPath(p[0], p[1], p[2]), nil
	}

	return nil, cerr.ActorPathError
}
