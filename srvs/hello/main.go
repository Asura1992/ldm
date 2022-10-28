package main

import (
	"ldm/common/constant"
	"ldm/common/dao"
	"ldm/common/protos/hello"
	"ldm/initalize"
	"ldm/srvs/hello/srv"
	"log"
)

func main() {
	//初始化配置
	initalize.InitGlobalConfig()
	//初始化redis
	initalize.InitRedis()
	//初始化数据库
	initalize.InitMysql()
	//初始化服务
	service, jaegerCloser, err := initalize.InitService(constant.API_HELLO_SRV)
	if err != nil {
		log.Fatal(err)
	}
	defer jaegerCloser.Close()
	if err := hello.RegisterHelloHandler(service.Server(), srv.NewHelloImplImpl(service.Client(), dao.Db)); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal("运行错误：", err)
	}
}
