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
	ActorConvertPathError   int32 = 26 // convert to path error
	ActorMarshalError       int32 = 27 // marshal arg error
	ActorUnmarshalError     int32 = 28 // unmarshal arg error
	ActorCallFail           int32 = 29 // actor call fail
	ActorSourceEqualTarget  int32 = 30 // source equal target
	ActorPublishRemoteError int32 = 31 // actor publish remote error
	ActorChildIDNotFound    int32 = 32 // actor child id not found

)

var (
	Error                    int32 = 1   // error
	PIDError                 int32 = 100 // pid错误
	LoginError               int32 = 201 // 登录异常
	AccountAuthFail          int32 = 202 // 帐号授权失败
	AccountBindFail          int32 = 203 // 帐号绑定失败
	AccountTokenValidateFail int32 = 204 // token验证失败
	AccountNameIsExist       int32 = 205 // 帐号已存在
	AccountRegisterError     int32 = 206 //
	AccountGetFail           int32 = 207 //
	PlayerDenyLogin          int32 = 301 // 玩家禁止登录
	PlayerDuplicateLogin     int32 = 302 // 玩家重复登录
	PlayerNameExist          int32 = 303 // 玩家角色名已存在
	PlayerCreateFail         int32 = 304 // 玩家创建角色失败
	PlayerNotLogin           int32 = 305 // 玩家未登录
	PlayerIdError            int32 = 306 // 玩家id错误
)
