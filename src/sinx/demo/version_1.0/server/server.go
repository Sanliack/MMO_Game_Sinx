package main

import (
	"fmt"
	"mmm/sinx/siface"
	"mmm/sinx/simodel"
	"strconv"
)

type s22 struct {
	simodel.RouteModel
}

type s11 struct {
	simodel.RouteModel
}

func (s *s22) Handle(req siface.RequestFace) {
	recmsg := req.GetMsg()
	fmt.Printf("s22 : server收到%d号消息:%s\n", recmsg.GetMsgID(), string(recmsg.GetData()))
	req.GetMsg()

	pack := simodel.NewPackMsgModel()
	msg, _ := pack.PackMsgByOther(99, []byte("s22_server返回消息id99。"))
	_, err := req.GetConn().GetTCPConn().Write(msg)
	if err != nil {
		fmt.Println("server send error", err)
	}
	return
}

func (s *s11) Handle(req siface.RequestFace) {
	req.GetConn().GetConnAddrMap().SetAddr("aaaa", 666666)
	recmsg := req.GetMsg()
	fmt.Printf("s11 : server收到%d号消息:%s\n", recmsg.GetMsgID(), string(recmsg.GetData()))
	req.GetMsg()

	pack := simodel.NewPackMsgModel()

	ggg := strconv.Itoa(req.GetConn().GetConnAddrMap().GetAddr("aaaa").(int))
	msg, _ := pack.PackMsgByOther(90, []byte("s11_server返回消息id99。ConnAddrtest:"+ggg))
	_, err := req.GetConn().GetTCPConn().Write(msg)
	if err != nil {
		fmt.Println("server send error", err)
	}
	return
}

func main() {
	sever := simodel.NewSinxServer()
	sever.RegisterHookFuncOnStart(func(c siface.ConnFace) {
		fmt.Printf("conn ID:%d Ip:%s connStartFunc被调用 Conn启动。。。\n", c.GetConnID(), c.RemoteADDR().String())
	})
	sever.RegisterHookFuncOnStop(func(c siface.ConnFace) {
		fmt.Printf("conn ID:%d Ip:%s connStopFunc被调用 Conn结束啦。。。\n", c.GetConnID(), c.RemoteADDR().String())
	})
	sever.AddRoute(11, &s11{})
	sever.AddRoute(22, &s22{})
	sever.Server()
}
