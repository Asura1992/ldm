package main

import (
	"ldm/common/constant"
	"ldm/common/dao"
	"ldm/common/protos/liveroom"
	"ldm/initalize"
	"ldm/srvs/liveroom/repos"
	"ldm/srvs/liveroom/srv"
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
	service, err := initalize.InitService(constant.API_LIVEROOM_SRV)
	if err != nil {
		log.Fatal(err)
	}
	repo := repos.NewLiveroomModel(dao.Db)
	if err := liveroom.RegisterLiveroomMicroHandler(service.Server(), srv.NewLiveroomImpl(service.Client(), repo)); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal("运行错误：", err)
	}
}
