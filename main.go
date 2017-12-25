package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"seed/config"
	"seed/service"
	"syscall"
)

var (
	// GitHash 编译脚本填写,推荐采用短hash
	GitHash = "unknown"
	// CompileTime 编译脚本填写,推荐精确到秒
	CompileTime = "unknown"
)

func main() {
	// confFile := flag.String("config", "seed.conf", "配置文件路径")
	ptrGrpcPort := flag.String("grpc", "8820", "GRPC端口")
	ptrRestfulPort := flag.String("restful", "8821", "RESTFUL端口")
	flag.Parse()

	service.Start(service.Option{
		GrpcPort:    *ptrGrpcPort,
		RestfulPort: *ptrRestfulPort,
	})

	fmt.Println(
		"GitHash", GitHash,
		"CompileTime", CompileTime,
		"ConfigVar", config.ConfigArg,
		"GrpcPort", *ptrRestfulPort,
		"RestfulPort", *ptrRestfulPort,
	)

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGTERM)

	s := <-signalChan

	fmt.Println("recvice signal:" + s.String())
}
