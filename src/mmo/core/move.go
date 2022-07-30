package core

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	__ "mmm/mmo/pb"
	"mmm/sinx/siface"
	"mmm/sinx/simodel"
)

type Move struct {
	simodel.RouteModel
}

func (m *Move) Handle(req siface.RequestFace) {
	msg := &__.Position{}
	err := proto.Unmarshal(req.GetMsg().GetData(), msg)
	if err != nil {
		fmt.Println("move accept error", err)
		return
	}
	pid := req.GetConn().GetConnAddrMap().GetAddr("pid").(int)
	player := Worldmanager.GetPlayerByPid(pid)
	fmt.Printf("player%d 移动到了(X:%f,Y:%f,Z:%f,V:%f)\n", pid, msg.X, msg.Y, msg.Z, msg.V)
	player.SyncMove(msg.X, msg.Y, msg.Z, msg.V)

}

func NewMove() *Move {
	return &Move{}
}
