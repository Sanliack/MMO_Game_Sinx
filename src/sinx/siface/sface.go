package siface

type Server interface {
	Start()
	Stop()
	Server()
	AddRoute(uint32, RouteFace)
	StartWorkerPool()
	GetMsgHandle() MsgHandleFace
	GetConnManager() ConnManagerFace
	GetHookFunc() HookFuncFace
	RegisterHookFuncOnStop(func(ConnFace))
	RegisterHookFuncOnStart(func(ConnFace))
}
