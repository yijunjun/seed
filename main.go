package main

import (
	"flag"
	"fmt"
	"seed/config"
	"seed/service"
)

var (
	// GitHash 编译脚本填写,推荐采用短hash
	GitHash = "unknown"
	// CompileTime 编译脚本填写,推荐精确到秒
	CompileTime = "unknown"
)

func main() {
	// confFile := flag.String("config", "seed.conf", "配置文件路径")
	ptrPort := flag.String("port", "1120", "监听端口")
	flag.Parse()
	service.Start(*ptrPort)
	fmt.Println("GitHash", GitHash, "CompileTime", CompileTime, "ConfigVar", config.ConfigArg)
}
