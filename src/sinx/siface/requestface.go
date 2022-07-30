package siface

type RequestFace interface {
	GetConn() ConnFace
	GetMsg() MessageFace
}
