package simple

import (
	. "github.com/po2656233/superplace/const"
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	cactor "github.com/po2656233/superplace/net/actor"
	cproto "github.com/po2656233/superplace/net/proto"
	"google.golang.org/protobuf/proto"
)

type ActorBase struct {
	cactor.Base
	session *cproto.Session
}

// SetSession 设置节点信息
func (p *ActorBase) SetSession(session *cproto.Session) {
	p.session = session
}

// SendMsg 发送消息
func (p *ActorBase) SendMsg(message proto.Message) {
	if onProtoFunc != nil {
		mid, data, err := onProtoFunc(message)
		if err != nil {
			clog.Errorf("[sid = %s,uid = %d] SendMsg fail. [mid = %d, message = %+v]",
				p.session.Sid,
				p.session.Uid,
				mid,
				message,
			)
			return
		}
		rsp := &cproto.PomeloResponse{
			Sid:  p.session.Sid,
			Mid:  mid,
			Data: data,
		}

		p.Call(p.session.AgentPath, ResponseFuncName, rsp)
	} else {
		clog.Panicf("Did you forget to set SetParseProtoFunc???")
	}
}

func (p *ActorBase) SendTo(sid string, message proto.Message) {
	if onProtoFunc != nil {
		mid, data, err := onProtoFunc(message)
		if err != nil {
			clog.Errorf("[sid = %s,uid = %d] SendMsg fail. [mid = %d, message = %+v]",
				sid,
				p.session.Uid,
				mid,
				message,
			)
			return
		}
		rsp := &cproto.PomeloResponse{
			Sid:  sid,
			Mid:  mid,
			Data: data,
		}

		p.Call(p.session.AgentPath, ResponseFuncName, rsp)
	} else {
		clog.Panicf("Did you forget to set SetParseProtoFunc???")
	}
}

// Response 响应
func (p *ActorBase) Response(session *cproto.Session, v interface{}) {
	Response(p, session, session.Mid, v)
}

// Feedback 反馈
func (p *ActorBase) Feedback(v interface{}) {
	if onProtoFunc != nil {
		p.SendMsg(v.(proto.Message))
		return
	}
	data, err := p.App().Serializer().Marshal(v)
	if err != nil {
		clog.Warnf("[Feedback] Marshal error. v = %+v", v)
		return
	}
	rsp := &cproto.PomeloResponse{
		Sid:  p.session.Sid,
		Mid:  p.session.Mid,
		Data: data,
	}
	p.Call(p.session.AgentPath, ResponseFuncName, rsp)
}

func SendTo(iActor cfacade.IActor, session *cproto.Session, v interface{}) {
	if onProtoFunc != nil {
		mid, data, err := onProtoFunc(v.(proto.Message))
		if err != nil {
			clog.Errorf("[sid = %s,uid = %d] SendMsg fail. [mid = %d, message = %+v]",
				session.Sid,
				session.Uid,
				mid,
				v,
			)
			return
		}
		rsp := &cproto.SimpleRequest{
			Sid:  session.Sid,
			Uid:  session.Uid,
			Mid:  mid,
			Data: data,
		}
		iActor.Call(session.AgentPath, RequestFuncName, rsp)
		return
	}

	data, err := iActor.App().Serializer().Marshal(v)
	if err != nil {
		clog.Warnf("[Response] Marshal error. v = %+v", v)
		return
	}

	rsp := &cproto.SimpleRequest{
		Sid:  session.Sid,
		Uid:  session.Uid,
		Mid:  session.Mid,
		Data: data,
	}

	iActor.Call(session.AgentPath, RequestFuncName, rsp)
}

// Response 响应
func Response(iActor cfacade.IActor, session *cproto.Session, mid uint32, v interface{}) {
	data, err := iActor.App().Serializer().Marshal(v)
	if err != nil {
		clog.Warnf("[Response] Marshal error. v = %+v", v)
		return
	}

	rsp := &cproto.PomeloResponse{
		Sid:  session.Sid,
		Mid:  mid,
		Data: data,
	}

	iActor.Call(session.AgentPath, ResponseFuncName, rsp)
}
