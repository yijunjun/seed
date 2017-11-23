package main

import (
	"fmt"
	"seed/config"
)

var (
	// GitHash 编译脚本填写,推荐采用短hash
	GitHash = "unknown"
	// CompileTime 编译脚本填写,推荐精确到秒
	CompileTime = "unknown"
)

func main() {
	fmt.Println("GitHash", GitHash, "CompileTime", CompileTime, "ConfigVar", config.ConfigArg)
}
