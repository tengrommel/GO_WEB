package main

import (
	"net/rpc"
	"net"
	"log"
)

type MathService struct {

}

type AddRequest struct {
	M int
	N int
}

type AddResponse struct {
	Result int
}

func (m *MathService)Add(request *AddRequest, reply *AddResponse) error {
	reply.Result = request.M + request.N
	return nil
}

func main() {
	mathService := new(MathService)
	rpc.Register(mathService)
	l, err := net.Listen("tcp", ":8021")
	if err !=nil{
		log.Fatal(err)
	}
	rpc.Accept(l)
}
