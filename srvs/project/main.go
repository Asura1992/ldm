package main

import (
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	grpcsvr "github.com/go-micro/plugins/v4/server/grpc"
	"ldm/common/config"
	"ldm/common/constant"
	"ldm/common/protos/project"
	"ldm/initalize"
	"ldm/srvs/project/impl"
	"log"
	"strings"
	"time"
)

func main(){
	//初始化配置
	if err := initalize.InitGlobalConfig();err != nil{
		log.Fatal(err)
	}
	cfg := config.GlobalConfig
	service := micro.NewService(
		micro.Server(grpcsvr.NewServer()),//这个要加上，不然grpc网关路由调用不会等待返回
		micro.Name(constant.API_PROJECT_SRV),
		micro.RegisterInterval(time.Second * 10),
		micro.RegisterTTL(time.Second * 5),
		micro.Registry(etcd.NewRegistry(registry.Addrs(strings.Split(cfg.Etcd.Address,",")...))),
		)
	service.Init()
	if err := project.RegisterProjectHandler(service.Server(),impl.NewProjectImpl(service.Client()));err != nil{
		log.Fatal(err)
	}
	if err := service.Run();err != nil{
		log.Fatal("运行错误：",err)
	}
}
