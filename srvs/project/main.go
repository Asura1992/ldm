package main

import (
	"ldm/common/constant"
	"ldm/common/protos/project"
	"ldm/initalize"
	"ldm/srvs/project/impl"
	"github.com/go-micro/plugins/v4/client/grpc"
	"log"
)

func main(){
	//初始化配置
	initalize.InitGlobalConfig()
	//初始化数据库
	initalize.InitMysql()
	service :=initalize.InitService(constant.API_PROJECT_SRV)
	if err := project.RegisterProjectHandler(service.Server(),impl.NewProjectImpl(grpc.NewClient()));err != nil{
		log.Fatal(err)
	}
	if err := service.Run();err != nil{
		log.Fatal("运行错误：",err)
	}
}
