package main

import (
	"fmt"
	"sinx/siface"
	"sinx/simodel"
)

type sr struct {
	simodel.RouteModel
}

func (s *sr) PreHandle(req siface.RequestFace) {
	go func() {
		for {
			fmt.Println("wait msg!!!")
			box := make([]byte, 512)
			n, err := req.GetConn().GetTCPConn().Read(box)
			if err != nil {
				fmt.Println("server Read error", err)
				return
			}
			pack := simodel.NewPackMsgModel()
			recmsg, err := pack.UnPackMsg(box[:n])
			if err != nil {
				fmt.Println("server UnpackMsg error", err)
				return
			}
			fmt.Printf("server收到ID:%d,长度为:%d,内容是:%s\n", recmsg.GetMsgID(), recmsg.GetMsgLen(), string(recmsg.GetData()))
		}
	}()
}

func (s *sr) Handle(req siface.RequestFace) {
	content := []byte("Sinx Version_0.5,server发送给client-----亚瑟摩根")
	msg := simodel.NewMessageModel(66, uint32(len(content)), content)
	pack := simodel.NewPackMsgModel()
	databyte, err := pack.PackMsg(msg)
	if err != nil {
		fmt.Println("pack_error", err)
		return
	}
	_, err = req.GetConn().GetTCPConn().Write(databyte)
	if err != nil {
		fmt.Println("handle error", err)
	}

}

func main() {
	sever := simodel.NewSinxServer()
	sever.AddRoute(&sr{})
	sever.Server()
}
