package simodel

import (
	"fmt"
	"mmm/sinx/siface"
	"sync"
)

type ConnManagerModel struct {
	ConnMaps map[uint32]siface.ConnFace
	connLock sync.RWMutex
}

func (c *ConnManagerModel) AddConn(conn siface.ConnFace) {
	c.connLock.Lock()
	c.ConnMaps[conn.GetConnID()] = conn
	c.connLock.Unlock()

}

func (c *ConnManagerModel) RemoveConn(u uint32) {
	c.connLock.Lock()
	_, ok := c.ConnMaps[u]
	if !ok {
		fmt.Printf("请求删除ID=%d的Conn不存在\n", u)
		return
	}
	delete(c.ConnMaps, u)
	c.connLock.Unlock()
}

func (c *ConnManagerModel) ClearConns() {
	c.ConnMaps = map[uint32]siface.ConnFace{}
}

func (c *ConnManagerModel) GetConnNums() uint32 {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	return uint32(len(c.ConnMaps))
}

func NewConnManagerModel() siface.ConnManagerFace {
	return &ConnManagerModel{
		ConnMaps: make(map[uint32]siface.ConnFace, 20),
	}
}
