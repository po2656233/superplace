package simple

import (
	"encoding/binary"
	"google.golang.org/protobuf/proto"
	"net"
	"time"

	"github.com/nats-io/nuid"
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	cactor "github.com/po2656233/superplace/net/actor"
	cproto "github.com/po2656233/superplace/net/proto"
	"go.uber.org/zap/zapcore"
)

type (
	actor struct {
		cactor.Base
		agentActorID   string
		connectors     []cfacade.IConnector
		onNewAgentFunc OnNewAgentFunc
		parseProtoFunc ParseProtoFunc // --Unnecessary--
	}

	OnNewAgentFunc func(newAgent *Agent)
	ParseProtoFunc func(message proto.Message) (uint32, []byte, error)
)

func NewActor(agentActorID string) *actor {
	if agentActorID == "" {
		panic("agentActorID is empty.")
	}

	parser := &actor{
		agentActorID: agentActorID,
		connectors:   make([]cfacade.IConnector, 0),
	}

	return parser
}

// OnInit Actor初始化前触发该函数
func (p *actor) OnInit() {
	p.Remote().Register(p.response)
}

func (p *actor) Load(app cfacade.IApplication) {
	if len(p.connectors) < 1 {
		panic("Connectors is nil. Please call the AddConnector(...) method add IConnector.")
	}

	//  Create agent actor
	if _, err := app.ActorSystem().CreateActor(p.agentActorID, p); err != nil {
		clog.Panicf("Create agent actor fail. err = %+v", err)
	}

	for _, connector := range p.connectors {
		connector.OnConnect(p.defaultOnConnectFunc)
		go connector.Start() // start connector!
	}
}

func (p *actor) AddConnector(connector cfacade.IConnector) {
	p.connectors = append(p.connectors, connector)
}

func (p *actor) Connectors() []cfacade.IConnector {
	return p.connectors
}

func (p *actor) AddNodeRoute(mid uint32, nodeRoute *NodeRoute) {
	AddNodeRoute(mid, nodeRoute)
}

// defaultOnConnectFunc 创建新连接时，通过当前agentActor创建child agent actor
func (p *actor) defaultOnConnectFunc(conn net.Conn) {
	session := &cproto.Session{
		Sid:       nuid.Next(),
		AgentPath: p.Path().String(),
		Data:      map[string]string{},
	}

	agent := NewAgent(p.App(), conn, session)

	if p.onNewAgentFunc != nil {
		p.onNewAgentFunc(&agent)
	}

	BindSID(&agent)
	agent.Run()
}

func (p *actor) SetOnNewAgent(fn OnNewAgentFunc) {
	p.onNewAgentFunc = fn
}

func (p *actor) SetHeartbeatTime(t time.Duration) {
	SetHeartbeatTime(t)
}

func (p *actor) SetWriteBacklog(backlog int) {
	SetWriteBacklog(backlog)
}

func (p *actor) SetEndian(e binary.ByteOrder) {
	SetEndian(e)
}
func (p *actor) SetMsgMaxLen(size uint32) {
	msgMaxLen = size
}

func (*actor) SetOnDataRoute(fn DataRouteFunc) {
	if fn != nil {
		onDataRouteFunc = fn
	}
}

func (p *actor) response(rsp *cproto.PomeloResponse) {
	agent, found := GetAgent(rsp.Sid)
	if !found {
		if clog.PrintLevel(zapcore.DebugLevel) {
			clog.Debugf("[response] Not found agent. [rsp = %+v]", rsp)
		}
		return
	}

	agent.Response(rsp.Mid, rsp.Data)
}
