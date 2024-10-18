package simple

import (
	. "github.com/po2656233/superplace/const"
	code2 "github.com/po2656233/superplace/const/code"
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
		mid, _, err := onProtoFunc(message)
		if err != nil {
			clog.Errorf("[sid = %s,uid = %d] SendMsg fail. [mid = %d, message = %+v]",
				p.session.Sid,
				p.session.Uid,
				mid,
				message,
			)
			return
		}
		data, _ := p.App().Serializer().Marshal(message)
		rsp := &cproto.PomeloResponse{
			Sid:  p.session.Sid,
			Mid:  mid,
			Data: data,
		}

		if code := p.Call(p.session.AgentPath, ResponseFuncName, rsp); code != code2.OK {
			clog.Warnf("[Response] Call error. code: %v", code)
		}
	} else {
		clog.Panicf("Did you forget to set SetParseProtoFunc???")
	}
}

func (p *ActorBase) SendTo(sid string, message proto.Message) {
	if onProtoFunc != nil {
		mid, _, err := onProtoFunc(message)
		if err != nil {
			clog.Errorf("[sid = %s,uid = %d] SendMsg fail. [mid = %d, message = %+v]",
				sid,
				p.session.Uid,
				mid,
				message,
			)
			return
		}
		data, err := p.App().Serializer().Marshal(message)
		rsp := &cproto.PomeloResponse{
			Sid:  sid,
			Mid:  mid,
			Data: data,
		}

		if code := p.Call(p.session.AgentPath, ResponseFuncName, rsp); code != code2.OK {
			clog.Warnf("[Response] Call error. code: %v", code)
		}
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
	if code := p.Call(p.session.AgentPath, ResponseFuncName, rsp); code != code2.OK {
		clog.Warnf("[Response] Call error. code: %v", code)
	}
}

func SendTo(iActor cfacade.IActor, session *cproto.Session, v interface{}) {
	if onProtoFunc != nil {
		mid, _, err := onProtoFunc(v.(proto.Message))
		if err != nil {
			clog.Errorf("[sid = %s,uid = %d] SendMsg fail. [mid = %d, message = %+v]",
				session.Sid,
				session.Uid,
				mid,
				v,
			)
			return
		}
		data, err := iActor.App().Serializer().Marshal(v)
		rsp := &cproto.SimpleRequest{
			Sid:  session.Sid,
			Uid:  session.Uid,
			Mid:  mid,
			Data: data,
		}
		if code := iActor.Call(session.AgentPath, RequestFuncName, rsp); code != code2.OK {
			clog.Warnf("[Response] Call error. code: %v", code)
		}
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

	if code := iActor.Call(session.AgentPath, RequestFuncName, rsp); code != code2.OK {
		clog.Warnf("[Response] Call error. code: %v", code)
	}
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

	if code := iActor.Call(session.AgentPath, ResponseFuncName, rsp); code != code2.OK {
		clog.Warnf("[Response] Call error. code: %v", code)
	}
}
