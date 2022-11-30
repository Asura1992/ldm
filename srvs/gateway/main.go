package main

import (
	"ldm/initalize"
	"ldm/srvs/gateway/srv"
	"log"
)

func main() {
	//初始化配置
	initalize.InitGlobalConfig()
	if err := srv.InitGateway(); err != nil {
		log.Fatal(err)
	}
}
