package simple

import (
	face "superplace/facade"
	clog "superplace/logger"
	cproto "superplace/net/proto"
)

var (
	nodeRouteMap    = map[uint32]*NodeRoute{}
	onDataRouteFunc = DefaultDataRoute
)

type (
	NodeRoute struct {
		NodeType string
		ActorID  string
		FuncName string
	}

	DataRouteFunc func(agent *Agent, msg *Message, route *NodeRoute)
)

func AddNodeRoute(mid uint32, nodeRoute *NodeRoute) {
	if nodeRoute == nil {
		return
	}

	nodeRouteMap[mid] = nodeRoute
}

func GetNodeRoute(mid uint32) (*NodeRoute, bool) {
	routeActor, found := nodeRouteMap[mid]
	return routeActor, found
}

func DefaultDataRoute(agent *Agent, msg *Message, route *NodeRoute) {
	session := agent.session
	session.Mid = msg.MID

	// current node
	if agent.NodeType() == route.NodeType {
		targetPath := face.NewChildPath(agent.NodeId(), route.ActorID, session.Sid)
		LocalDataRoute(agent, session, msg, route, targetPath)
		return
	}

	if !session.IsBind() {
		clog.Warnf("[sid = %s,uid = %d] Session is not bind with UID. failed to forward message.[route = %+v]",
			agent.SID(),
			agent.UID(),
			route,
		)
		return
	}

	member, found := agent.Discovery().Random(route.NodeType)
	if !found {
		return
	}

	targetPath := face.NewPath(member.GetNodeId(), route.ActorID)
	ClusterLocalDataRoute(agent, session, msg, route, member.GetNodeId(), targetPath)
}

func LocalDataRoute(agent *Agent, session *cproto.Session, msg *Message, nodeRoute *NodeRoute, targetPath string) {
	message := face.GetMessage()
	message.Source = session.AgentPath
	message.Target = targetPath
	message.FuncName = nodeRoute.FuncName
	message.Session = session
	message.Args = msg.Data

	agent.ActorSystem().PostLocal(message)
}

func ClusterLocalDataRoute(agent *Agent, session *cproto.Session, msg *Message, nodeRoute *NodeRoute, nodeID, targetPath string) error {
	clusterPacket := cproto.GetClusterPacket()
	clusterPacket.SourcePath = session.AgentPath
	clusterPacket.TargetPath = targetPath
	clusterPacket.FuncName = nodeRoute.FuncName
	clusterPacket.Session = session   // agent session
	clusterPacket.ArgBytes = msg.Data // packet -> message -> base

	return agent.Cluster().PublishLocal(nodeID, clusterPacket)
}