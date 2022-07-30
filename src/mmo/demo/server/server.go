package main

import (
	"fmt"
	"mmm/mmo/core"
	"mmm/sinx/siface"
	"mmm/sinx/simodel"
)

func HookFuncStart(conn siface.ConnFace) {
	newplayer := core.NewAPlayer(conn)
	newplayer.SendPidToClient()
	newplayer.BroadCastStartAddr()
	newplayer.SendAddrToOther()
	conn.GetConnAddrMap().SetAddr("pid", newplayer.Pid)
	core.Worldmanager.AddPlayer(newplayer)
	fmt.Printf("Player成功连接上服务器，ID:=%d,当前服务器拥有玩家数：%d。\n", newplayer.Pid, core.Worldmanager.PlayerCount)
}

func HookFuncStop(conn siface.ConnFace) {
	pid := conn.GetConnAddrMap().GetAddr("pid").(int)
	player := core.Worldmanager.GetPlayerByPid(pid)
	player.Offline()
	fmt.Printf("player%d offline success\n", pid)
	fmt.Println(core.Worldmanager.Players, core.Worldmanager.GetAllPlayer())
}

func main() {
	s := simodel.NewSinxServer()
	s.RegisterHookFuncOnStart(HookFuncStart)
	s.RegisterHookFuncOnStop(HookFuncStop)
	s.AddRoute(2, core.NewWorldChat())
	s.AddRoute(3, core.NewMove())
	s.Server()
}
