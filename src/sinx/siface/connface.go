package siface

import "net"

type ConnFace interface {
	Start()
	Stop()
	GetTCPConn() *net.TCPConn
	GetConnID() uint32
	RemoteADDR() net.Addr
	Send(RequestFace) error
	StartWriter()
	GetMsgHandle() MsgHandleFace
	GetMsgChan() chan RequestFace
	GetServer() Server
	SendBufMsg(int, []byte) error
	GetConnAddrMap() SetConnAddrFace
}

type HandleFunc func(*net.TCPConn, []byte) error
