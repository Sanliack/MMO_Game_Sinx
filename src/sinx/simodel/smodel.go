package simodel

import (
	"fmt"
	"mmm/sinx/config"
	"mmm/sinx/siface"
	"net"
)

type SinxServer struct {
	Name         string
	NetType      string
	IP           string
	Port         int
	MaxTranSize  int
	MsgHandleMap siface.MsgHandleFace
	ConnManager  siface.ConnManagerFace
	HookFunc     siface.HookFuncFace
}

func (s *SinxServer) Start() {
	tcpaddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("获取tcpAddr出错", err)
	}
	listen, err := net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		fmt.Println("获取sinxServer-listen出错", err)
	}
	var connid uint32 = 0
	s.StartWorkerPool()
	fmt.Printf("%s成功启动,地址为%s:%d,Worker数量为%d，WorkerChan长度为%d\n", s.Name, s.IP, s.Port, s.GetMsgHandle().GetMsgWorkerPoolNum(), config.SinxConfig.MsgQueueLen)
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("conn出错", err)
			continue
		}
		if s.ConnManager.GetConnNums() >= uint32(config.SinxConfig.MaxConnections) {
			fmt.Println("服务器已到达允许的最大连接数,,已拒绝一个新的连接，目标地址:"+conn.RemoteAddr().String()+";目前连接数%d", s.ConnManager.GetConnNums())
			_ = conn.Close()
			continue
		}
		newConn := NewConnModel(conn, connid, s.MsgHandleMap, s)
		connid++
		go newConn.Start()
		s.ConnManager.AddConn(newConn)
		fmt.Printf("%s开启了一条Conn协程ID=%d;并添加到了connmanager中,对应地址为%v\n", s.Name, newConn.ConnID, newConn.Conn.RemoteAddr())
		fmt.Printf("server最大支持连接数:%d;当前conn数%d\n", config.SinxConfig.MaxConnections, s.ConnManager.GetConnNums())
		fmt.Println(config.SinxConfig)
	}
}

func (s *SinxServer) Stop() {
	defer fmt.Printf("server：Name:%s 地址：%s:%d 服务器正常退出 资源已释放。\n", s.Name, s.IP, s.Port)
	s.GetConnManager().ClearConns()
	s.GetMsgHandle().Stop()
}

func (s *SinxServer) Server() {
	defer s.Stop()
	s.Start()
	select {}
}

func (s *SinxServer) GetHookFunc() siface.HookFuncFace {
	return s.HookFunc
}

func (s *SinxServer) RegisterHookFuncOnStart(fu func(siface.ConnFace)) {
	s.GetHookFunc().RegisterOnConnStart(fu)
	fmt.Printf("server：%s,Ip: %s:%d 成功注册ConnStartHookfunc\n", s.Name, s.IP, s.Port)

}

func (s *SinxServer) RegisterHookFuncOnStop(fu func(siface.ConnFace)) {
	s.GetHookFunc().RegisterOnConnStop(fu)
	fmt.Printf("server：%s,Ip: %s:%d 成功注册ConnStopHookfunc\n", s.Name, s.IP, s.Port)
}

func (s *SinxServer) AddRoute(id uint32, route siface.RouteFace) {
	s.MsgHandleMap.AddRoute(id, route)
}

func (s *SinxServer) StartWorkerPool() {
	s.MsgHandleMap.StartWorkerPool()
}

func (s *SinxServer) GetMsgHandle() siface.MsgHandleFace {
	return s.MsgHandleMap
}

func (s *SinxServer) GetConnManager() siface.ConnManagerFace {
	return s.ConnManager
}
func NewSinxServer() *SinxServer {
	connmanager := NewConnManagerModel()
	return &SinxServer{
		Name:         config.SinxConfig.Name,
		NetType:      "TCP",
		IP:           config.SinxConfig.Host,
		Port:         config.SinxConfig.Port,
		MaxTranSize:  config.SinxConfig.MaxTranSize,
		MsgHandleMap: NewMsgHandleModel(),
		ConnManager:  connmanager,
		HookFunc:     NewHookFuncModel(),
	}
}
