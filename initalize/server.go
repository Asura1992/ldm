package initalize

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	grpcsvr "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"ldm/common/config"
	"strings"
	"time"
)

//初始化服务
func InitService(srvName string, WrapHandler ...server.HandlerWrapper) micro.Service {
	microOpt := []micro.Option{
		micro.Server(grpcsvr.NewServer()), //这个要加上，不然grpc网关路由调用不会等待返回
		micro.Name(srvName),
		micro.RegisterInterval(time.Second * 5),
		micro.RegisterTTL(time.Second * 10),
		micro.Registry(etcd.NewRegistry(registry.Addrs(strings.Split(config.GlobalConfig.Etcd.Address, ",")...))),
	}
	if len(WrapHandler) > 0 {
		microOpt = append(microOpt, micro.WrapHandler(WrapHandler[0]))
	}
	service := micro.NewService(microOpt...)
	service.Init()
	return service
}

type ShopInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

//拦截器
func WrapHandle(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		//获取头信息
		md, b := metadata.FromContext(ctx)
		if !b {
			return errors.New("metadata not found")
		}
		//TODO 拦截设置,可根据token验证是否合理请求
		fmt.Println("拦截器获取元数据:", md)
		//通过验证调用服务
		err := fn(ctx, req, rsp)
		if err != nil {
			return err
		}
		return nil
	}
}
