package main

import (
	"github.com/go-micro/plugins/v4/registry/etcd"
	grpcsvr "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"ldm/common/constant"
	"ldm/common/protos/hello"
	"ldm/initalize"
	"ldm/srvs/hello/impl"
	"log"
	"strings"
	"time"
)

func main(){
	//初始化配置
	if err := initalize.InitGlobalConfig();err != nil{
		log.Fatal(err)
	}
	cfg := initalize.GlobalConfig
	service := micro.NewService(
		micro.Server(grpcsvr.NewServer()),//这个要加上，不然grpc网关路由调用不会等待返回
		micro.Name(constant.API_HELLO_SRV),
		micro.RegisterInterval(time.Second * 15),
		micro.RegisterTTL(time.Second * 30),
		micro.Registry(etcd.NewRegistry(registry.Addrs(strings.Split(cfg.Etcd.Address,",")...))),
		)
	service.Init()
	if err := hello.RegisterHelloHandler(service.Server(),impl.NewHelloImplImpl(service.Client()));err != nil{
		log.Fatal(err)
	}
	if err := service.Run();err != nil{
		log.Fatal("运行错误：",err)
	}
}
