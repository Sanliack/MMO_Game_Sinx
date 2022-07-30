package simodel

import (
	"mmm/sinx/siface"
)

type RequestModel struct {
	Conn *ConnModel
	Msg  siface.MessageFace
}

func (r *RequestModel) GetConn() siface.ConnFace {
	return r.Conn
}

func (r *RequestModel) GetMsg() siface.MessageFace {
	return r.Msg
}

func NewRequestModel(conn ConnModel, msg siface.MessageFace) *RequestModel {
	return &RequestModel{
		Conn: &conn,
		Msg:  msg,
	}
}
