package simodel

import (
	"fmt"
	"mmm/sinx/config"
	"mmm/sinx/siface"
)

type MsgHandleModel struct {
	IdRouteMap       map[uint32]siface.RouteFace
	MsgQueue         []chan siface.RequestFace
	MaxWorkerPoolNum uint32
}

func (r *MsgHandleModel) AddRoute(id uint32, route siface.RouteFace) {
	_, ok := r.IdRouteMap[id]
	if ok == true {
		fmt.Printf("ID为%d的对应路由已经存在，无需重复添加。", id)
		return
	}
	r.IdRouteMap[id] = route
	fmt.Printf("server: 成功添加Id为%d的对应路由\n", id)
}

func (r *MsgHandleModel) Handle(req siface.RequestFace) {
	route, ok := r.GetRoute(req.GetMsg().GetMsgID())
	if ok {
		route.PreHandle(req)
		route.Handle(req)
		route.AftHandle(req)
	} else {
		fmt.Printf("client发送了并未添加的路由！！！\n")
		return
	}

}

func (r *MsgHandleModel) Stop() {
	r.MsgQueue = []chan siface.RequestFace{}
	r.IdRouteMap = map[uint32]siface.RouteFace{}
}

func (r *MsgHandleModel) GetRoute(id uint32) (siface.RouteFace, bool) {
	rface, ok := r.IdRouteMap[id]
	return rface, ok
}

func (r *MsgHandleModel) StartWorkerPool() {
	for i, v := range r.MsgQueue {
		go r.startOneWorker(v)
		fmt.Printf("成功启动%d号worker协程\n", i)
	}
}

func (r *MsgHandleModel) GetMsgQueue() []chan siface.RequestFace {
	return r.MsgQueue
}

func (r *MsgHandleModel) GetMsgQueueByAvg(id uint32) chan siface.RequestFace {
	afterid := id % r.MaxWorkerPoolNum
	return r.MsgQueue[afterid]
}

func (r *MsgHandleModel) startOneWorker(ch chan siface.RequestFace) {
	defer close(ch)
	for {
		select {
		case req := <-ch:
			//req.GetConn().GetMsgChan() <- req
			req.GetConn().GetMsgHandle().Handle(req)
		}
	}
}

func (r *MsgHandleModel) GetMsgWorkerPoolNum() uint32 {
	return r.MaxWorkerPoolNum
}

func NewMsgHandleModel() *MsgHandleModel {
	msgq := make([]chan siface.RequestFace, config.SinxConfig.MsgQueueNum)
	for i := 0; i < config.SinxConfig.MsgQueueNum; i++ {
		msgq[i] = make(chan siface.RequestFace, config.SinxConfig.MsgQueueLen)
	}
	return &MsgHandleModel{
		IdRouteMap:       map[uint32]siface.RouteFace{},
		MsgQueue:         msgq,
		MaxWorkerPoolNum: uint32(config.SinxConfig.MsgQueueNum),
	}
}
