package siface

type SetConnAddrFace interface {
	SetAddr(string, interface{})
	GetAddr(string) interface{}
}
