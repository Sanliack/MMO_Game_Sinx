package main

import (
	"fmt"
	"net"
	"sinx/simodel"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("dial error ", err)
		return
	}
	packmodel := simodel.NewPackMsgModel()
	msg1 := simodel.NewMessageModel()
	msg2 := simodel.NewMessageModel()
	msg1.SetMsgID(14)
	msg1.SetData([]byte("我带你打"))
	msg1.SetMsgLen(uint32(len(msg1.GetData())))

	msg2.SetMsgID(99)
	msg2.SetData([]byte("啊哈哈哈哈哈哈adasff"))
	msg2.SetMsgLen(uint32(len(msg2.GetData())))

	msg1byte, _ := packmodel.PackMsg(msg1)
	msg2byte, _ := packmodel.PackMsg(msg2)

	msg1byte = append(msg1byte, msg2byte...)
	//_, err = conn.Write(msg2byte)
	//fmt.Println(string(msg1byte))
	_, err = conn.Write(msg1byte)

	fmt.Println(msg1, msg2)
	if err != nil {
		fmt.Println("error", err)
	}
}
