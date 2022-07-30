package siface

type ConnManagerFace interface {
	AddConn(ConnFace)
	RemoveConn(uint32)
	ClearConns()
	GetConnNums() uint32
}
