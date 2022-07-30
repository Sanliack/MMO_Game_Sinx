package core

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"math/rand"
	__ "mmm/mmo/pb"
	"mmm/sinx/siface"
)

var playno = 1

// 1 pid , 200 addr
type Player struct {
	Pid  int
	X    float32
	Y    float32
	Z    float32
	V    float32
	Conn siface.ConnFace
}

func NewAPlayer(conn siface.ConnFace) *Player {
	player := &Player{
		playno,
		float32(160 + rand.Intn(20)),
		0,
		float32(160 + rand.Intn(20)),
		0,
		conn,
	}
	playno++
	return player
}

func (p *Player) SendMsg(msgID int, msg proto.Message) {
	if p.Conn == nil {
		fmt.Println("Conn连接未建立")
		return
	}
	bufmsg, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("player message TO []byte error", err)
		return
	}
	err = p.Conn.SendBufMsg(msgID, bufmsg)
	if err != nil {
		fmt.Println("Send player message error", err)
	}
}

func (p *Player) SendPidToClient() {
	player := &__.Player{
		Pid: int64(p.Pid),
	}
	p.SendMsg(1, player)
}

func (p *Player) Talk(content string) {
	players := Worldmanager.GetAllPlayer()
	realmsg := &__.BroadCast_Content{
		Content: content,
	}
	msg := &__.BroadCast{
		Pid:     int64(p.Pid),
		MsgType: 1,
		Msg:     realmsg,
	}
	for _, player := range players {
		player.SendMsg(200, msg)
	}
}

func (p *Player) BroadCastStartAddr() {
	pos := &__.Position{
		X: p.X,
		Y: p.Y,
		Z: p.Z,
		V: p.V}
	playpos := &__.BroadCast_P{
		P: pos,
	}
	msg := &__.BroadCast{
		Pid:     int64(p.Pid),
		MsgType: int64(2),
		Msg:     playpos,
	}
	p.SendMsg(200, msg)
}

func (p *Player) SendAddrToOther() {
	Posmsg := &__.Position{
		X: p.X,
		Y: p.Y,
		Z: p.Z,
		V: p.V,
	}
	pos := &__.BroadCast_P{
		P: Posmsg,
	}
	mmm := &__.BroadCast{
		Pid:     int64(p.Pid),
		MsgType: 2,
		Msg:     pos,
	}
	allplayer := Worldmanager.AOI.GetSurroundGridPlayersIDsByXY(int(p.X), int(p.Z))
	fmt.Println("玩家列表:", allplayer)
	var playerandpos []*__.PlayerAndPos
	for _, v := range allplayer {
		remoteplayer := Worldmanager.GetPlayerByPid(v)
		remoteplayer.SendMsg(200, mmm)
		playerandpos = append(playerandpos, &__.PlayerAndPos{
			Pid: int32(v),
			Pos: &__.Position{
				X: remoteplayer.X,
				Y: remoteplayer.Y,
				Z: remoteplayer.Z,
				V: remoteplayer.V,
			},
		})
	}

	other := &__.SyncPlays{Players: playerandpos}
	p.SendMsg(202, other)
}

func (p *Player) SyncMove(x, y, z, v float32) {
	p.X = x
	p.Y = y
	p.Z = z
	p.V = v

	broadmsg := &__.BroadCast{
		Pid:     int64(p.Pid),
		MsgType: 4,
		Msg: &__.BroadCast_P{
			P: &__.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	pids := Worldmanager.AOI.GetSurroundGridPlayersIDsByXY(int(p.X), int(p.Z))
	for _, v := range pids {
		if v == p.Pid {
			continue
		}
		Worldmanager.GetPlayerByPid(v).SendMsg(200, broadmsg)
	}
}

func (p *Player) Offline() {
	nearp := Worldmanager.AOI.GetSurroundGridPlayersIDsByXY(int(p.X), int(p.Z))
	msg := &__.Player{
		Pid: int64(p.Pid),
	}
	for _, v := range nearp {
		if v == p.Pid {
			continue
		}
		Worldmanager.GetPlayerByPid(v).SendMsg(201, msg)
	}
	//Worldmanager.AOI.RemovePlaysByXY(p.Pid, int(p.X), int(p.Z))
	Worldmanager.RemovePlayerByPid(p.Pid)

}
