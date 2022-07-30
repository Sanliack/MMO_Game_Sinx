package main

import (
	"fmt"
	"net"
)

func main() {
	Hehe()
	conn, err := net.Dial("tcp", "127.0.0.1:33366")
	if err != nil {
		fmt.Println("client err")
	}
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(buf[:n]))
	}

}

func Hehe() {
	var aa = make(map[int]int, 20)
	fmt.Println(len(aa))
	aa[1] = 1
	aa[2] = 1
	aa[3] = 1
	aa[4] = 1
	aa[5] = 1
	fmt.Println(len(aa))
	delete(aa, 5)
	fmt.Println(len(aa))

}
