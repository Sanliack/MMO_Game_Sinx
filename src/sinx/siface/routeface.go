package siface

type RouteFace interface {
	PreHandle(RequestFace)
	Handle(RequestFace)
	AftHandle(RequestFace)
}
