package core

import (
	"mmm/sinx/siface"
	"mmm/sinx/simodel"
)

type WorldChat struct {
	simodel.RouteModel
}

func (w *WorldChat) Handle(req siface.RequestFace) {
	con := req.GetMsg().GetData()
	pid := req.GetConn().GetConnAddrMap().GetAddr("pid").(int)
	player := Worldmanager.GetPlayerByPid(pid)
	player.Talk(string(con))

}

func NewWorldChat() *WorldChat {
	return &WorldChat{}
}
