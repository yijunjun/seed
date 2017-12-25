package service

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	grpc "google.golang.org/grpc"
)

type service struct {
	Grpc    func(*grpc.Server)
	Restful func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) (err error)
}

var gServices []*service

func reg(rs *service) {
	gServices = append(gServices, rs)
}

// Option 启动服务选项
type Option struct {
	GrpcPort    string
	RestfulPort string
}

// Start 启动所有服务
func Start(opt Option) error {
	if opt.GrpcPort == opt.RestfulPort {
		return fmt.Errorf(
			"grpc(%v) == restful(%v)",
			opt.GrpcPort,
			opt.RestfulPort,
		)
	}

	if opt.GrpcPort != "" {
		return errors.New(
			"grpc is empty",
		)
	}

	port := ":" + opt.GrpcPort
	l, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	gs := grpc.NewServer()
	for _, s := range gServices {
		if s.Grpc != nil {
			s.Grpc(gs)
		}
	}

	if opt.RestfulPort != "" {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}

		for _, s := range gServices {
			if s.Restful != nil {
				if err := s.Restful(ctx, mux, port, opts); err != nil {
					return err
				}
			}
		}

		go http.ListenAndServe(":"+opt.RestfulPort, mux)
	}

	go gs.Serve(l)
	return nil
}
