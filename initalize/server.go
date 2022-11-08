package initalize

import (
	"context"
	"errors"
	"fmt"
	grpc_cli "github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/go-micro/plugins/v4/server/grpc"
	microOpentracing "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"io"
	"ldm/common/config"
	"ldm/utils/jaeger"
	"strings"
	"time"
)

//初始化服务
func InitService(srvName string, authWrapHandler ...server.HandlerWrapper) (micro.Service, io.Closer, error) {
	cfg := config.GlobalConfig
	//etcd集群地址
	etcdAddrArray := strings.Split(cfg.Etcd.Address, ",")
	microOpt := []micro.Option{
		micro.Server(grpc.NewServer()), //这个要加上，不然grpc网关路由调用不会等待返回
		micro.Name(srvName),
		micro.RegisterInterval(time.Second * 5), //每5秒重新注册
		micro.RegisterTTL(time.Second * 10),     //10秒过期
		micro.Version(time.Now().Format("2006-01-02 15:04:05")),
		micro.Client(grpc_cli.NewClient()), //client要用grpc的
		micro.Registry(etcd.NewRegistry(registry.Addrs(etcdAddrArray...))),
	}
	//服务端链路追踪
	microOpt = append(microOpt, micro.WrapHandler(microOpentracing.NewHandlerWrapper(opentracing.GlobalTracer())))
	//拦截器
	if len(authWrapHandler) > 0 {
		microOpt = append(microOpt, micro.WrapHandler(authWrapHandler[0]))
	}
	service := micro.NewService(microOpt...)
	service.Init()
	//链路追踪
	_, jaegerCloser, err := jaeger.NewJaegerTracer(srvName, cfg.Jaeger.JaegerTracerAddr)
	if err != nil {
		return nil, jaegerCloser, err
	}
	return service, jaegerCloser, nil
}

type ShopInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

//拦截器
func AuthWrapHandle(fn server.HandlerFunc) server.HandlerFunc {
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
