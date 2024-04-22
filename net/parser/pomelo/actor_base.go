package pomelo

import (
	face "superplace/facade"
	clog "superplace/logger"
	cactor "superplace/net/actor"
	cproto "superplace/net/proto"
)

const (
	ResponseFuncName = "response"
	PushFuncName     = "push"
	KickFuncName     = "kick"
	BroadcastName    = "broadcast"
)

type ActorBase struct {
	cactor.Base
}

func (p *ActorBase) Response(session *cproto.Session, v interface{}) {
	Response(p, session.AgentPath, session.Sid, session.Mid, v)
}

func (p *ActorBase) ResponseCode(session *cproto.Session, statusCode int32) {
	ResponseCode(p, session.AgentPath, session.Sid, session.Mid, statusCode)
}

func (p *ActorBase) Push(session *cproto.Session, route string, v interface{}) {
	Push(p, session.AgentPath, session.Sid, route, v)
}

func (p *ActorBase) Kick(session *cproto.Session, reason interface{}, closed bool) {
	Kick(p, session.AgentPath, session.Sid, reason, closed)
}

func (p *ActorBase) Broadcast(agentPath string, uidList []int64, allUID bool, route string, v interface{}) {
	Broadcast(p, agentPath, uidList, allUID, route, v)
}

func Response(iActor face.IActor, agentPath, sid string, mid uint32, v interface{}) {
	data, err := iActor.App().Serializer().Marshal(v)
	if err != nil {
		clog.Warnf("[Response] Marshal error. v = %+v", v)
		return
	}

	rsp := &cproto.PomeloResponse{
		Sid:  sid,
		Mid:  mid,
		Data: data,
	}

	iActor.Call(agentPath, ResponseFuncName, rsp)
}

func ResponseCode(iActor face.IActor, agentPath, sid string, mid uint32, statusCode int32) {
	rsp := &cproto.PomeloResponse{
		Sid:  sid,
		Mid:  mid,
		Code: statusCode,
	}

	iActor.Call(agentPath, ResponseFuncName, rsp)
}

func Push(iActor face.IActor, agentPath, sid, route string, v interface{}) {
	if route == "" {
		clog.Warn("[Push] route value error.")
		return
	}

	data, err := iActor.App().Serializer().Marshal(v)
	if err != nil {
		clog.Warnf("[Push] Marshal error. route =%s, v = %+v", route, v)
		return
	}

	rsp := &cproto.PomeloPush{
		Sid:   sid,
		Route: route,
		Data:  data,
	}

	iActor.Call(agentPath, PushFuncName, rsp)
}

func Kick(iActor face.IActor, agentPath, sid string, reason interface{}, closed bool) {
	data, err := iActor.App().Serializer().Marshal(reason)
	if err != nil {
		clog.Warnf("[Kick] Marshal error. reason = %+v", reason)
		return
	}

	rsp := &cproto.PomeloKick{
		Sid:    sid,
		Reason: data,
		Close:  closed,
	}

	iActor.Call(agentPath, KickFuncName, rsp)
}

func Broadcast(iActor face.IActor, agentPath string, uidList []int64, allUID bool, route string, v interface{}) {
	if !allUID && len(uidList) < 1 {
		clog.Warn("[Broadcast] uidList value error.")
		return
	}

	if route == "" {
		clog.Warn("[Broadcast] route value error.")
		return
	}

	data, err := iActor.App().Serializer().Marshal(v)
	if err != nil {
		clog.Warnf("[Kick] Marshal error. v = %+v", v)
		return
	}

	rsp := &cproto.PomeloBroadcastPush{
		UidList: uidList,
		AllUID:  allUID,
		Route:   route,
		Data:    data,
	}

	iActor.Call(agentPath, BroadcastName, rsp)
}
