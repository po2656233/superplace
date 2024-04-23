package pomelo

import (
	face "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	pmessage "github.com/po2656233/superplace/net/parser/pomelo/message"
	cproto "github.com/po2656233/superplace/net/proto"
)

// DefaultDataRoute 默认的消息路由
func DefaultDataRoute(agent *Agent, route *pmessage.Route, msg *pmessage.Message) {
	session := BuildSession(agent, msg)

	// current node
	if agent.NodeType() == route.NodeType() {
		targetPath := face.NewChildPath(agent.NodeId(), route.HandleName(), session.Sid)
		LocalDataRoute(agent, session, route, msg, targetPath)
		return
	}

	if !session.IsBind() {
		clog.Warnf("[sid = %s,uid = %d] Session is not bind with UID. failed to forward message.[route = %s]",
			agent.SID(),
			agent.UID(),
			msg.Route,
		)
		return
	}

	member, found := agent.Discovery().Random(route.NodeType())
	if !found {
		return
	}

	targetPath := face.NewPath(member.GetNodeId(), route.HandleName())
	err := ClusterLocalDataRoute(agent, session, route, msg, member.GetNodeId(), targetPath)
	if err != nil {
		clog.Warnf("[sid = %s,uid = %d,route = %s] cluster local base error. err= %v",
			agent.SID(),
			agent.UID(),
			msg.Route,
			err,
		)
	}
}

func LocalDataRoute(agent *Agent, session *cproto.Session, route *pmessage.Route, msg *pmessage.Message, targetPath string) {
	message := face.GetMessage()
	message.Source = session.AgentPath
	message.Target = targetPath
	message.FuncName = route.Method()
	message.Session = session
	message.Args = msg.Data

	agent.ActorSystem().PostLocal(message)
}

func ClusterLocalDataRoute(agent *Agent, session *cproto.Session, route *pmessage.Route, msg *pmessage.Message, nodeID, targetPath string) error {
	clusterPacket := cproto.GetClusterPacket()
	clusterPacket.SourcePath = session.AgentPath
	clusterPacket.TargetPath = targetPath
	clusterPacket.FuncName = route.Method()
	clusterPacket.Session = session   // agent session
	clusterPacket.ArgBytes = msg.Data // packet -> message -> base

	return agent.Cluster().PublishLocal(nodeID, clusterPacket)
}

func BuildSession(agent *Agent, msg *pmessage.Message) *cproto.Session {
	agent.session.Mid = uint32(msg.ID)
	return agent.session
}
