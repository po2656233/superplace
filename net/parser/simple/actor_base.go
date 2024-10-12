package simple

import (
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

func (p *ActorBase) Response(session *cproto.Session, v interface{}) {
	Response(p, session, session.Mid, v)
}

func (p *ActorBase) ResponseX(session *cproto.Session, mid uint32, v interface{}) {
	Response(p, session, mid, v)
}

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
