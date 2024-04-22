package cherryDiscovery

import (
	"math/rand"
	cerr "superplace/logger/error"
	"sync"

	cprofile "superplace/config"
	cslice "superplace/extend/slice"
	face "superplace/facade"
	clog "superplace/logger"
	cproto "superplace/net/proto"
)

// DiscoveryDefault 默认方式，通过读取profile文件的节点信息
//
// 该类型发现服务仅用于开发测试使用，直接读取profile.json->node配置
type DiscoveryDefault struct {
	memberMap        sync.Map // key:nodeId,value:face.IMember
	onAddListener    []face.MemberListener
	onRemoveListener []face.MemberListener
}

func (n *DiscoveryDefault) PreInit() {
	n.memberMap = sync.Map{}
}

func (n *DiscoveryDefault) Load(_ face.IApplication) {
	// load node info from config file
	nodeConfig := cprofile.GetConfig("node")
	if nodeConfig.LastError() != nil {
		clog.Error("`node` property not found in config file.")
		return
	}

	for _, nodeType := range nodeConfig.Keys() {
		typeJson := nodeConfig.Get(nodeType)
		for i := 0; i < typeJson.Size(); i++ {
			item := typeJson.Get(i)

			nodeId := item.Get("node_id").ToString()
			if nodeId == "" {
				clog.Errorf("nodeId is empty in nodeType = %s", nodeType)
				break
			}

			if _, found := n.GetMember(nodeId); found {
				clog.Errorf("nodeType = %s, nodeId = %s, duplicate nodeId", nodeType, nodeId)
				break
			}

			member := &cproto.Member{
				NodeId:   nodeId,
				NodeType: nodeType,
				Address:  item.Get("rpc_address").ToString(),
				Settings: make(map[string]string),
			}

			settings := item.Get("__settings__")
			for _, key := range settings.Keys() {
				member.Settings[key] = settings.Get(key).ToString()
			}

			n.memberMap.Store(member.NodeId, member)
		}
	}
}

func (n *DiscoveryDefault) Name() string {
	return "default"
}

func (n *DiscoveryDefault) Map() map[string]face.IMember {
	memberMap := map[string]face.IMember{}

	n.memberMap.Range(func(key, value any) bool {
		if member, ok := value.(face.IMember); ok {
			memberMap[member.GetNodeId()] = member
		}
		return true
	})

	return memberMap
}

func (n *DiscoveryDefault) ListByType(nodeType string, filterNodeId ...string) []face.IMember {
	var memberList []face.IMember

	n.memberMap.Range(func(key, value any) bool {
		member := value.(face.IMember)
		if member.GetNodeType() == nodeType {
			if _, ok := cslice.StringIn(member.GetNodeId(), filterNodeId); !ok {
				memberList = append(memberList, member)
			}
		}

		return true
	})

	return memberList
}

func (n *DiscoveryDefault) Random(nodeType string) (face.IMember, bool) {
	memberList := n.ListByType(nodeType)
	memberLen := len(memberList)

	if memberLen < 1 {
		return nil, false
	}

	if memberLen == 1 {
		return memberList[0], true
	}

	return memberList[rand.Intn(len(memberList))], true
}

func (n *DiscoveryDefault) GetType(nodeId string) (nodeType string, err error) {
	member, found := n.GetMember(nodeId)
	if !found {
		return "", cerr.Errorf("nodeId = %s not found.", nodeId)
	}
	return member.GetNodeType(), nil
}

func (n *DiscoveryDefault) GetMember(nodeId string) (face.IMember, bool) {
	if nodeId == "" {
		return nil, false
	}

	value, found := n.memberMap.Load(nodeId)
	if !found {
		return nil, false
	}

	return value.(face.IMember), found
}

func (n *DiscoveryDefault) AddMember(member face.IMember) {
	_, loaded := n.memberMap.LoadOrStore(member.GetNodeId(), member)
	if loaded {
		clog.Warnf("duplicate nodeId. [nodeType = %s], [nodeId = %s], [address = %s]",
			member.GetNodeType(),
			member.GetNodeId(),
			member.GetAddress(),
		)
		return
	}

	for _, listener := range n.onAddListener {
		listener(member)
	}

	clog.Debugf("addMember new member. [member = %s]", member)
}

func (n *DiscoveryDefault) RemoveMember(nodeId string) {
	value, loaded := n.memberMap.LoadAndDelete(nodeId)
	if loaded {
		member := value.(face.IMember)
		clog.Debugf("remove member. [member = %s]", member)

		for _, listener := range n.onRemoveListener {
			listener(member)
		}
	}
}

func (n *DiscoveryDefault) OnAddMember(listener face.MemberListener) {
	if listener == nil {
		return
	}
	n.onAddListener = append(n.onAddListener, listener)
}

func (n *DiscoveryDefault) OnRemoveMember(listener face.MemberListener) {
	if listener == nil {
		return
	}
	n.onRemoveListener = append(n.onRemoveListener, listener)
}

func (n *DiscoveryDefault) Stop() {

}