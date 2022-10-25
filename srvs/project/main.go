package main

import (
	"ldm/common/constant"
	"ldm/common/dao"
	"ldm/common/protos/project"
	"ldm/initalize"
	"ldm/srvs/project/srv"
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
	service := initalize.InitService(constant.API_PROJECT_SRV, initalize.WrapHandle)
	if err := project.RegisterProjectHandler(service.Server(), srv.NewProjectImpl(service.Client(), dao.Db)); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal("运行错误：", err)
	}
}
