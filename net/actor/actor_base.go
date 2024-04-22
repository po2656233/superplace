package cherryActor

import (
	face "superplace/facade"
)

type Base struct {
	Actor
}

func (p *Base) load(a Actor) {
	p.Actor = a
}

func (p *Base) AliasID() string {
	return ""
}

// OnInit Actor初始化前触发该函数
func (*Base) OnInit() {
}

// OnStop Actor停止前触发该函数
func (*Base) OnStop() {
}

// OnLocalReceived Actor收到Local消息时触发该函数
func (*Base) OnLocalReceived(_ *face.Message) (next bool, invoke bool) {
	next = true
	invoke = false
	return
}

// OnRemoteReceived Actor收到Remote消息时触发该函数
func (*Base) OnRemoteReceived(_ *face.Message) (next bool, invoke bool) {
	next = true
	invoke = false
	return
}

// OnFindChild 寻找子Actor时触发该函数.开发者可以自定义创建子Actor
func (*Base) OnFindChild(_ *face.Message) (face.IActor, bool) {
	return nil, false
}

func (p *Base) NewPath(nodeID, actorID interface{}) string {
	return face.NewPath(nodeID, actorID)
}

func (p *Base) NewNodePath(actorID interface{}) string {
	return face.NewPath(p.path.NodeID, actorID)
}

func (p *Base) NewChildPath(actorID, childID interface{}) string {
	return face.NewChildPath(p.path.NodeID, actorID, childID)
}

func (p *Base) NewMyChildPath(childID interface{}) string {
	return face.NewChildPath(p.path.NodeID, p.path.ActorID, childID)
}
