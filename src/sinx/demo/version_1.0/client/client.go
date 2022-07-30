package main

import (
	"fmt"
	"io"
	"mmm/sinx/simodel"
	"net"
	"time"
)

//func main() {
//	for i := 0; i < 10; i++ {
//		go man()
//	}
//	select {}
//}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:33366")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("client", err)
		return
	}
	var sendid uint32 = 11
	go ListenConn(conn)
	ll := 0

	for {
		if ll%3 == 0 {
			sendid = 22
		} else {
			sendid = 11
		}
		ll++
		pack := simodel.NewPackMsgModel()
		msg, err := pack.PackMsgByOther(sendid, []byte("client向server发送了条消息，id为11号。finish。"))
		if err != nil {
			fmt.Println("client have error", err)
			continue
		}
		_, err = conn.Write(msg)
		if err != nil && err == io.EOF {
			fmt.Printf("server断开了conn（可能原因：server已到设定最大连接上线）")
			return
		} else if err != nil {
			fmt.Println("server断开了conn（可能原因：server已到设定最大连接上线）", err)
			return
		}
		time.Sleep(time.Second * 1)
	}
}

func ListenConn(conn *net.TCPConn) {
	pack := simodel.NewPackMsgModel()
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Printf("server断开了conn（可能原因：server已到设定最大连接上线）")
			return
		} else if err != nil {

			fmt.Println("server断开了conn（可能原因：server已到设定最大连接上线）", err)
			return
		}
		msg, err := pack.UnPackMsg(buf[:n])
		if err != nil {
			fmt.Println("client unpack 出现了错误", err)
			return
		}
		fmt.Printf("client接收到id为:%d-的新消息，内容为%s\n", msg.GetMsgID(), msg.GetData())
	}
}
