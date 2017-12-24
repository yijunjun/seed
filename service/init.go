package service

import (
	"net"

	grpc "google.golang.org/grpc"
)

type regServiceHandler func(*grpc.Server)

var gServiceHandlers []regServiceHandler

func regService(handler regServiceHandler) {
	gServiceHandlers = append(gServiceHandlers, handler)
}

// Start 启动所有服务
func Start(port string) error {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	gs := grpc.NewServer()
	for _, h := range gServiceHandlers {
		h(gs)
	}
	return gs.Serve(l)
}
