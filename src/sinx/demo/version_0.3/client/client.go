package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:33366")
	if err != nil {
		fmt.Println("client", err)
		return
	}
	_, err = conn.Write([]byte("俺是个怀念收购艾格尼丝懊悔噶阿松哟i啊说过啊"))
	if err != nil {
		fmt.Println("client_write", err)
		return
	}

	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client接收err", err)
			continue
		}
		fmt.Println(string(buf[:n]))

	}
}
