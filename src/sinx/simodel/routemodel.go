package simodel

import "mmm/sinx/siface"

type RouteModel struct{}

func (r *RouteModel) PreHandle(siface.RequestFace) {}

func (r *RouteModel) Handle(siface.RequestFace) {}

func (r *RouteModel) AftHandle(siface.RequestFace) {}

func NewRouteModel() *RouteModel {
	return &RouteModel{}
}
