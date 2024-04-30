package code

const (
	OK                    int32 = 0  // is ok
	SessionUIDNotBind     int32 = 10 // session uid not bind
	DiscoveryNotFoundNode int32 = 11 // discovery not fond node id
	NodeRequestError      int32 = 12 // node request error

	RPCNetError           int32 = 20 // rpc net error
	RPCUnmarshalError     int32 = 21 // rpc base unmarshal error
	RPCMarshalError       int32 = 22 // rpc base marshal error
	RPCRemoteExecuteError int32 = 23 // rpc remote method executor error

	ActorPathIsNil          int32 = 24 // actor target path is nil
	ActorFuncNameError      int32 = 25 // actor function name is error
	ActorConvertPathError   int32 = 26 // actor convert to path error
	ActorMarshalError       int32 = 27 // actor marshal arg error
	ActorUnmarshalError     int32 = 28 // actor unmarshal arg error
	ActorCallFail           int32 = 29 // actor call fail
	ActorSourceEqualTarget  int32 = 30 // actor source equal target
	ActorPublishRemoteError int32 = 31 // actor publish remote error
	ActorChildIDNotFound    int32 = 32 // actor child id not found
)

var CodeTxt = map[int32]string{
	OK:                      "is ok",
	SessionUIDNotBind:       "session uid not bind",
	DiscoveryNotFoundNode:   "discovery not fond node id",
	NodeRequestError:        "node request error",
	RPCNetError:             "rpc net error",
	RPCUnmarshalError:       "rpc base unmarshal error",
	RPCMarshalError:         "rpc base marshal error",
	RPCRemoteExecuteError:   "rpc remote method executor error",
	ActorPathIsNil:          "actor target path is nil",
	ActorFuncNameError:      "actor function name is error",
	ActorConvertPathError:   "convert to path error",
	ActorMarshalError:       "actor marshal arg error",
	ActorUnmarshalError:     "actor unmarshal arg error",
	ActorCallFail:           "actor call fail",
	ActorSourceEqualTarget:  "actor source equal target",
	ActorPublishRemoteError: "actor publish remote error",
	ActorChildIDNotFound:    "actor child id not found",
}
