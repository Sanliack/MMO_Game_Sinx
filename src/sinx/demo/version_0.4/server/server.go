package main

import (
	"fmt"
	"net"
	"sinx/simodel"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := listen.AcceptTCP()
		fmt.Println("conn 连接成功", conn)
		if err != nil {
			fmt.Println(err)
		}
		box := make([]byte, 4096)
		pack := simodel.NewPackMsgModel()
		n, err := conn.Read(box)
		if err != nil {
			fmt.Println(err)
		}
		go func(bb []byte) {
			for {
				fmt.Println("go_bb", bb)
				if len(bb) == 0 {
					fmt.Println("接收结束")
					return
				}
				msg, err := pack.UnPackMsg(bb)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("server接收消息第%d号信息，长度为%d,内容为:%v\n", msg.GetMsgID(), msg.GetMsgLen(), string(bb[pack.GetHeadLen():pack.GetHeadLen()+msg.GetMsgLen()]))
				bb = bb[pack.GetHeadLen()+msg.GetMsgLen():]
			}
		}(box[:n])

	}
}
