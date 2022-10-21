package initalize

import (
	"github.com/go-micro/plugins/v4/registry/etcd"
	grpcsvr "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"ldm/common/config"
	"strings"
	"time"
)
//初始化服务
func InitService(srvName string) micro.Service {
	service := micro.NewService(
		micro.Server(grpcsvr.NewServer()),//这个要加上，不然grpc网关路由调用不会等待返回
		micro.Name(srvName),
		micro.RegisterInterval(time.Second * 15),
		micro.RegisterTTL(time.Second * 30),
		micro.Registry(etcd.NewRegistry(registry.Addrs(strings.Split(config.GlobalConfig.Etcd.Address,",")...))),
	)
	service.Init()
	return service
}

