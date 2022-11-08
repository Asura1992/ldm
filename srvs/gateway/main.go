package main

import (
	"fmt"
	"ldm/initalize"
	"ldm/srvs/gateway/srv"
	"log"
)

func main() {
	//初始化配置
	initalize.InitGlobalConfig()
	fmt.Println()
	if err := srv.InitGateway(); err != nil {
		log.Fatal(err)
	}
}
