package main

import (
	"fmt"
	"net"
	"sinx/simodel"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:33366")
	if err != nil {
		fmt.Println("client", err)
		return
	}
	pack := simodel.NewPackMsgModel()
	go func() {
		for {
			buf := make([]byte, 512)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("client接收err", err)
				continue
			}
			recmsg, err := pack.UnPackMsg(buf[:n])
			if err != nil {
				fmt.Println("client unpack error", err)
				continue
			}
			fmt.Printf("client收到ID:%d,长度为:%d,内容是:%s\n", recmsg.GetMsgID(), recmsg.GetMsgLen(), string(recmsg.GetData()))
		}
	}()
	SendMsg(conn, "第一次发", 21)
	time.Sleep(time.Second * 1)
	SendMsg(conn, "333333333333", 22)
	//SendMsg(conn, "end最后", 23)
}

func SendMsg(conn net.Conn, i string, id uint32) {

	content := []byte("Sinx Version_0.5,连接测试文本QAQ===client发送给server。" + i)
	msg := simodel.NewMessageModel(id, uint32(len(content)), content)
	pack := simodel.NewPackMsgModel()
	databyte, err := pack.PackMsg(msg)
	if err != nil {
		fmt.Println("pack_error", err)
		return
	}
	_, err = conn.Write(databyte)
	if err != nil {
		fmt.Println("client_write", err)
		return
	}
}
