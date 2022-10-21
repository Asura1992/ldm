package main

import (
	"ldm/common/constant"
	"ldm/common/protos/hello"
	"ldm/initalize"
	"ldm/srvs/hello/impl"
	"log"
)

func main(){
	//初始化配置
	initalize.InitGlobalConfig()
	//初始化数据库
	initalize.InitMysql()
	service :=initalize.InitService(constant.API_HELLO_SRV)
	if err := hello.RegisterHelloHandler(service.Server(),impl.NewHelloImplImpl(service.Client()));err != nil{
		log.Fatal(err)
	}
	if err := service.Run();err != nil{
		log.Fatal("运行错误：",err)
	}
}
