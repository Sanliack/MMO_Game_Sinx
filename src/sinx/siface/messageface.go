package siface

type MessageFace interface {
	GetMsgID() uint32
	GetMsgLen() uint32
	GetData() []byte
	SetMsgLen(uint32)
	SetData([]byte)
	SetMsgID(uint32)
}
