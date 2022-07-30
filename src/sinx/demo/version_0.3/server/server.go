package main

import (
	"fmt"
	"sinx/siface"
	"sinx/simodel"
)

type sr struct {
	simodel.RouteModel
}

func (s *sr) PreHandle(req siface.RequestFace) {
	_, err := req.GetConn().GetTCPConn().Write([]byte("preHandle_test"))
	if err != nil {
		fmt.Println("prehandle error", err)
	}
}

func (s *sr) Handle(req siface.RequestFace) {
	_, err := req.GetConn().GetTCPConn().Write([]byte("Handle_test_test"))
	if err != nil {
		fmt.Println("handle error", err)
	}

}

func (s *sr) AftHandle(req siface.RequestFace) {
	_, err := req.GetConn().GetTCPConn().Write([]byte("aftHandle_test_test_test"))
	if err != nil {
		fmt.Println("afthandle error", err)
	}

}

func main() {
	sever := simodel.NewSinxServer("Version_0.3", 33366)
	sever.AddRoute(&sr{})
	sever.Server()
}
