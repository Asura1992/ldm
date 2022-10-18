package main

import (
	"github.com/go-micro/plugins/v4/registry/etcd"
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
		micro.Name(constant.API_HELLO_SRV),
		micro.RegisterInterval(time.Second * 15),
		micro.RegisterTTL(time.Second * 30),
		micro.Registry(etcd.NewRegistry(registry.Addrs(strings.Split(cfg.Etcd.Address,",")...))),
		)
	//优雅关闭服务
	service.Init()
	if err := hello.RegisterHelloHandler(service.Server(),impl.NewHelloImplImpl(service.Client()));err != nil{
		log.Fatal(err)
	}
	if err := service.Run();err != nil{
		log.Fatal("运行错误：",err)
	}
}
