package simodel

import (
	"fmt"
	"io"
	"mmm/sinx/config"
	"mmm/sinx/siface"
	"net"
)

type ConnModel struct {
	ConnID     uint32
	Conn       *net.TCPConn
	Status     int
	MsgHandle  siface.MsgHandleFace
	MsgChan    chan siface.RequestFace
	ConnServer siface.Server
	AddrMap    siface.SetConnAddrFace
}

func (c *ConnModel) Start() {
	//go c.StartWrite()
	c.GetServer().GetHookFunc().ConnStartFunc(c)
	c.StartReader()

}

func (c *ConnModel) GetMsgChan() chan siface.RequestFace {
	return c.MsgChan
}

func (c *ConnModel) StartReader() {
	for {
		box := make([]byte, config.SinxConfig.MaxTranSize)
		n, err := c.Conn.Read(box)
		if err != nil && err == io.EOF {
			fmt.Println("Conn连接传输结束")
			c.Stop()
			return
		} else if err != nil {
			fmt.Println("Conn start 从conn中读取[]byte出错", err)
			return
		}

		pack := NewPackMsgModel()
		msg, err := pack.UnPackMsg(box[:n])

		req := RequestModel{
			Conn: c,
			Msg:  msg,
		}

		c.GetMsgHandle().GetMsgQueueByAvg(req.GetMsg().GetMsgID()) <- &req
		fmt.Printf("COnn已将一条request放入%d的msgchan中\n", req.GetMsg().GetMsgID())
	}
}

func (c *ConnModel) GetMsgHandle() siface.MsgHandleFace {
	return c.MsgHandle
}

func (c *ConnModel) AddRequestToMsgChan(req siface.RequestFace) {
	c.GetMsgHandle()
}

func (c *ConnModel) StartWriter() {
	//for {
	//	select {
	//	case req := <-c.MsgChan:
	//		c.MsgHandle.Handle(req)
	//	case <-c.CloseChan:
	//		_ = c.Conn.Close()
	//		return
	//	}
	//}
}

func (c *ConnModel) GetServer() siface.Server {
	return c.ConnServer
}

func (c *ConnModel) Stop() {
	c.GetServer().GetHookFunc().ConnStopFunc(c)
	if c.Status != 0 {
		c.Status = 0
	}
	defer fmt.Printf("Conn：Id=%d成功关闭....\n", c.ConnID)
	_ = c.Conn.Close()
	c.GetServer().GetConnManager().RemoveConn(c.ConnID)
}

func (c *ConnModel) GetConnID() uint32 {
	return c.ConnID
}

func (c *ConnModel) GetTCPConn() *net.TCPConn {
	return c.Conn
}

func (c *ConnModel) RemoteADDR() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *ConnModel) Send(siface.RequestFace) error {
	return nil
}

func (c *ConnModel) SendBufMsg(msgid int, buf []byte) error {
	pack := NewPackMsgModel()
	buf, err := pack.PackMsgByOther(uint32(msgid), buf)
	if err != nil {
		fmt.Println("Conn Send PackMsgByOther error", err)
		return err
	}
	_, err = c.GetTCPConn().Write(buf)
	if err != nil {
		fmt.Println("Conn Send error", err)
		return err
	}
	return nil
}

func (c *ConnModel) GetConnAddrMap() siface.SetConnAddrFace {
	return c.AddrMap
}

func NewConnModel(conn *net.TCPConn, connid uint32, routemap siface.MsgHandleFace, server siface.Server) *ConnModel {
	return &ConnModel{
		ConnID:     connid,
		Conn:       conn,
		Status:     1,
		MsgHandle:  routemap,
		MsgChan:    make(chan siface.RequestFace),
		ConnServer: server,
		AddrMap:    NewSetAddrModel(),
	}
}
