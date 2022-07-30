package simodel

import "mmm/sinx/siface"

type HookFuncModel struct {
	OnConnStart func(siface.ConnFace)
	OnConnStop  func(siface.ConnFace)
}

func (h *HookFuncModel) RegisterOnConnStart(fu func(siface.ConnFace)) {
	h.OnConnStart = fu
}

func (h *HookFuncModel) RegisterOnConnStop(fu func(siface.ConnFace)) {
	h.OnConnStop = fu
}

func (h *HookFuncModel) ConnStartFunc(c siface.ConnFace) {
	h.OnConnStart(c)
}

func (h *HookFuncModel) ConnStopFunc(c siface.ConnFace) {
	h.OnConnStop(c)
}

func NewHookFuncModel() *HookFuncModel {
	return &HookFuncModel{
		OnConnStart: func(siface.ConnFace) {},
		OnConnStop:  func(siface.ConnFace) {},
	}
}
