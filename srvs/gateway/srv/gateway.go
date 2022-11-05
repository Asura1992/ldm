package srv

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go-micro.dev/v4/registry"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"ldm/common/config"
	"ldm/common/constant"
	"ldm/common/proto/protos/hello"
	"ldm/common/proto/protos/project"
	"log"
	"net/http"
	"strings"
	"time"
)

var gateWayMux = runtime.NewServeMux(
	//允许所有头信息
	runtime.WithIncomingHeaderMatcher(allowHeader),
	runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				Multiline:       true,
				UseProtoNames:   true,
				UseEnumNumbers:  true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true, //忽略传入非定义的字段
			},
		}))

//初始化网关
func InitGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	cfg := config.GlobalConfig
	reg := etcd.NewRegistry(registry.Addrs(strings.Split(cfg.Etcd.Address, ",")...))
	regSrvs, err := reg.ListServices(func(options *registry.ListOptions) {
		options.Context = ctx
	})
	if err != nil {
		log.Fatal(err)
	}
	//遍历所有etcd注册的服务
	for _, srv := range regSrvs {
		if err = registerEndpoint(ctx, *srv); err != nil {
			return err
		}
	}
	//启动swagger
	go initSwagger()
	//监听服务变化重新注册端点
	wathServiceChange(ctx, reg)
	//http监听服务启动
	listenAddr := fmt.Sprintf(":%d", config.GlobalConfig.HttpPort)
	connectTimeout := time.Second * time.Duration(config.GlobalConfig.HttpTimeout)
	return http.ListenAndServe(listenAddr, http.TimeoutHandler(gateWayMux, connectTimeout, http.ErrHandlerTimeout.Error()))
}

//注册端点
func registerEndpoint(ctx context.Context, srv registry.Service) (err error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	for _, node := range srv.Nodes {
		endpoint := flag.String(srv.Name+uuid.New().String(), node.Address, srv.Name)
		switch srv.Name {
		case constant.API_PROJECT_SRV: //项目服务
			err = project.RegisterProjectHandlerFromEndpoint(ctx, gateWayMux, *endpoint, opts)
		case constant.API_HELLO_SRV: //hello服务
			err = hello.RegisterHelloHandlerFromEndpoint(ctx, gateWayMux, *endpoint, opts)
		default:
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println(srv.Name + " 服务注册端点地址 " + node.Address)
	}
	return nil
}
