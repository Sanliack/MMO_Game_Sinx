package siface

type MsgHandleFace interface {
	AddRoute(uint32, RouteFace)
	Handle(RequestFace)
	GetRoute(uint32) (RouteFace, bool)
	StartWorkerPool()
	GetMsgQueue() []chan RequestFace
	GetMsgQueueByAvg(uint32) chan RequestFace
	GetMsgWorkerPoolNum() uint32
	Stop()
}
