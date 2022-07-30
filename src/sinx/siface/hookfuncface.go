package siface

type HookFuncFace interface {
	RegisterOnConnStart(func(ConnFace))
	RegisterOnConnStop(func(ConnFace))
	ConnStartFunc(ConnFace)
	ConnStopFunc(ConnFace)
}
