package cherryActor

import (
	"reflect"
	ccode "github/po2656233/superplace/const/code"
	cerror "github/po2656233/superplace/logger/error"

	"google.golang.org/protobuf/proto"

	creflect "github/po2656233/superplace/extend/reflect"
	cutils "github/po2656233/superplace/extend/utils"
	face "github/po2656233/superplace/facade"
	clog "github/po2656233/superplace/logger"
	cproto "github/po2656233/superplace/net/proto"
)

func InvokeLocalFunc(app face.IApplication, fi *creflect.FuncInfo, m *face.Message) {
	if app == nil {
		clog.Errorf("[InvokeLocalFunc] app is nil. [message = %+v]", m)
		return
	}

	EncodeLocalArgs(app, fi, m)

	values := make([]reflect.Value, 2)
	values[0] = reflect.ValueOf(m.Session) // session
	values[1] = reflect.ValueOf(m.Args)    // args
	fi.Value.Call(values)
}

func InvokeRemoteFunc(app face.IApplication, fi *creflect.FuncInfo, m *face.Message) {
	if app == nil {
		clog.Errorf("[InvokeRemoteFunc] app is nil. [message = %+v]", m)
		return
	}

	EncodeRemoteArgs(app, fi, m)

	values := make([]reflect.Value, fi.InArgsLen)
	if fi.InArgsLen > 0 {
		values[0] = reflect.ValueOf(m.Args) // args
	}

	if m.IsCluster {
		cutils.Try(func() {
			rets := fi.Value.Call(values)
			rspCode, rspData := retValue(app.Serializer(), rets)

			retResponse(m.ClusterReply, &cproto.Response{
				Code: rspCode,
				Data: rspData,
			})

		}, func(errString string) {
			retResponse(m.ClusterReply, &cproto.Response{
				Code: ccode.RPCRemoteExecuteError,
			})
			clog.Errorf("[InvokeRemoteFunc] invoke error. [message = %+v, err = %s]", m, errString)
		})
	} else {
		cutils.Try(func() {
			if m.ChanResult == nil {
				fi.Value.Call(values)
			} else {
				rets := fi.Value.Call(values)
				rspCode, rspData := retValue(app.Serializer(), rets)
				m.ChanResult <- &cproto.Response{
					Code: rspCode,
					Data: rspData,
				}
			}
		}, func(errString string) {
			if m.ChanResult != nil {
				m.ChanResult <- nil
			}

			clog.Errorf("[remote] invoke error.[source = %s, target = %s -> %s, funcType = %v, err = %+v]",
				m.Source,
				m.Target,
				m.FuncName,
				fi.InArgs,
				errString,
			)
		})
	}
}

func EncodeRemoteArgs(app face.IApplication, fi *creflect.FuncInfo, m *face.Message) error {
	if m.IsCluster {
		if fi.InArgsLen == 0 {
			return nil
		}

		return EncodeArgs(app, fi, 0, m)
	}

	return nil
}

func EncodeLocalArgs(app face.IApplication, fi *creflect.FuncInfo, m *face.Message) error {
	return EncodeArgs(app, fi, 1, m)
}

func EncodeArgs(app face.IApplication, fi *creflect.FuncInfo, index int, m *face.Message) error {
	argBytes, ok := m.Args.([]byte)
	if !ok {
		return cerror.Errorf("Encode args error.[source = %s, target = %s -> %s, funcType = %v]",
			m.Source,
			m.Target,
			m.FuncName,
			fi.InArgs,
		)
	}

	argValue := reflect.New(fi.InArgs[index].Elem()).Interface()
	err := app.Serializer().Unmarshal(argBytes, argValue)
	if err != nil {
		return cerror.Errorf("Encode args unmarshal error.[source = %s, target = %s -> %s, funcType = %v]",
			m.Source,
			m.Target,
			m.FuncName,
			fi.InArgs,
		)
	}

	m.Args = argValue

	return nil
}

func retValue(serializer face.ISerializer, rets []reflect.Value) (int32, []byte) {
	var (
		retsLen = len(rets)
		rspCode = ccode.OK
		rspData []byte
	)

	if retsLen == 1 {
		if val := rets[0].Interface(); val != nil {
			if c, ok := val.(int32); ok {
				rspCode = c
			}
		}
	} else if retsLen == 2 {
		if !rets[0].IsNil() {
			data, err := serializer.Marshal(rets[0].Interface())
			if err != nil {
				rspCode = ccode.RPCRemoteExecuteError
				clog.Warn(err)
			} else {
				rspData = data
			}
		}

		if val := rets[1].Interface(); val != nil {
			if c, ok := val.(int32); ok {
				rspCode = c
			}
		}
	}

	return rspCode, rspData
}

func retResponse(reply face.IRespond, rsp *cproto.Response) {
	if reply != nil {
		rspData, _ := proto.Marshal(rsp)
		err := reply.Respond(rspData)
		if err != nil {
			clog.Warn(err)
		}
	}
}
