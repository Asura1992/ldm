package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go-micro.dev/v4/registry"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"ldm/common/constant"
	"ldm/common/proto/gateway/protos/hello"
	"ldm/common/proto/gateway/protos/project"
	"ldm/initalize"
	"log"
	"net/http"
	"strings"
	"time"
)


func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
		MarshalOptions:protojson.MarshalOptions{
			UseProtoNames:   true,
			UseEnumNumbers:  true,
			EmitUnpopulated: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true, //忽略传入非定义的字段
		},
	}))
	if err := RegisterSrvEndpoint(ctx,mux);err != nil{
		return err
	}
	fmt.Println("gateway http listen on :" , initalize.GlobalConfig.HttpPort)
	return http.ListenAndServe(
		fmt.Sprintf(":%d", initalize.GlobalConfig.HttpPort),
		http.TimeoutHandler(mux,time.Minute * time.Duration(initalize.GlobalConfig.HttpTimeout),"api request timeout!!"),
		)
}
//注册服务端点供http调用
func RegisterSrvEndpoint(ctx context.Context,mux *runtime.ServeMux) error{
	cfg := initalize.GlobalConfig
	reg := etcd.NewRegistry(registry.Addrs(strings.Split(cfg.Etcd.Address,",")...))
	regSrvs,err := reg.ListServices(func(options *registry.ListOptions) {
		options.Context = ctx
	})
	if err != nil{
		log.Fatal(err)
	}
	opts := []grpc.DialOption{grpc.WithInsecure(),grpc.WithBlock()}
	//遍历所有etcd注册的服务
	for _,srv := range regSrvs{
		for _,node := range srv.Nodes{
			endpoint := flag.String(srv.Name,node.Address, srv.Name)
			switch srv.Name {
			case constant.API_PROJECT_SRV://项目服务
				err = project.RegisterProjectHandlerFromEndpoint(ctx, mux, *endpoint, opts)
			case constant.API_HELLO_SRV://hello服务
				err = hello.RegisterHelloHandlerFromEndpoint(ctx, mux, *endpoint, opts)
			}
			if err != nil {
				return err
			}
			fmt.Println(srv.Name + " 服务注册端点地址 "+node.Address)
		}
	}
	return nil
}

func main() {
	//初始化配置
	if err := initalize.InitGlobalConfig();err != nil{
		log.Fatal(err)
	}
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
