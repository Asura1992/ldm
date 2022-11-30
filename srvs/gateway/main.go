package main

import (
	"ldm/initalize"
	"ldm/srvs/gateway/code"
	"log"
)

func main() {
	//初始化配置
	initalize.InitGlobalConfig()
	if err := code.InitGateway(); err != nil {
		log.Fatal(err)
	}
}
