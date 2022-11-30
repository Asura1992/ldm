package main

import (
	"ldm/common/gateway/code"
	"ldm/initalize"
	"log"
)

func main() {
	//初始化配置
	initalize.InitGlobalConfig()
	if err := code.InitGateway(); err != nil {
		log.Fatal(err)
	}
}
