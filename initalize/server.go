package initalize

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/go-micro/plugins/v4/server/grpc"
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
		micro.Server(grpc.NewServer()), //这个要加上，不然grpc网关路由调用不会等待返回
		micro.Name(srvName),
		micro.RegisterInterval(time.Second * 5), //每5秒重新注册
		micro.RegisterTTL(time.Second * 10),     //10秒过期
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(registry.Addrs(strings.Split(config.GlobalConfig.Etcd.Address, ",")...))),
	}
	//拦截器
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
		fmt.Println("拦截器打印：", md)
		//TODO 拦截设置,可根据token验证是否合理请求,如果合法，则像下面把程序想要的数据写进上下文
		//ctx = metadata.Set(ctx, "ShopInfo", "6666666666666666666666666666")
		//通过验证调用服务
		if err := fn(ctx, req, rsp); err != nil {
			//TODO 写日志
			return errors.New(err.Error())
		}
		return nil
	}
}
