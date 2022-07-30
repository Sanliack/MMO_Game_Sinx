package siface

type PackFace interface {
	GetHeadLen() uint32
	PackMsg(MessageFace) ([]byte, error)
	PackMsgByOther(uint32, []byte) ([]byte, error)
	UnPackMsg([]byte) (MessageFace, error)
}
