package simple

import (
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	cproto "github.com/po2656233/superplace/net/proto"
)

var (
	nodeRouteMap    = map[uint32]*NodeRoute{}
	onDataRouteFunc = DefaultDataRoute
	onProtoFunc     ParseProtoFunc
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
func SetParseProtoFunc(protoFunc ParseProtoFunc) {
	onProtoFunc = protoFunc
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
		targetPath := cfacade.NewChildPath(agent.NodeId(), route.ActorID, session.Sid)
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

	targetPath := cfacade.NewPath(member.GetNodeId(), route.ActorID)
	if err := ClusterLocalDataRoute(agent, session, msg, route, member.GetNodeId(), targetPath); err != nil {
		clog.Errorf("[sid = %s,uid = %d] Session post message err:%v.[route = %+v]",
			agent.SID(),
			agent.UID(),
			err,
			route,
		)
	}
}

func LocalDataRoute(agent *Agent, session *cproto.Session, msg *Message, nodeRoute *NodeRoute, targetPath string) {
	message := cfacade.GetMessage()
	message.Source = session.AgentPath
	message.Target = targetPath
	message.FuncName = nodeRoute.FuncName
	message.Session = session
	message.Args = msg.Data

	agent.ActorSystem().PostLocal(&message)
}

func ClusterLocalDataRoute(agent *Agent, session *cproto.Session, msg *Message, nodeRoute *NodeRoute, nodeID, targetPath string) error {
	clusterPacket := cproto.GetClusterPacket()
	clusterPacket.SourcePath = session.AgentPath
	clusterPacket.TargetPath = targetPath
	clusterPacket.FuncName = nodeRoute.FuncName
	clusterPacket.Session = session   // agent session
	clusterPacket.ArgBytes = msg.Data // packet -> message -> data

	return agent.Cluster().PublishLocal(nodeID, clusterPacket)
}
