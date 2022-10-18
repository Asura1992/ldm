package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/google/uuid"
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
	//超时时间为1分钟
	ctx, cancel := context.WithTimeout(ctx,time.Minute)
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
	opts := []grpc.DialOption{grpc.WithInsecure(),grpc.WithBlock()}
	if err := RegisterSrvEndpoint(ctx,mux,opts);err != nil{
		return err
	}
	fmt.Println("gateway http listen on :" , initalize.GlobalConfig.HttpPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", initalize.GlobalConfig.HttpPort), mux)
}
//注册服务端点供http调用
func RegisterSrvEndpoint(ctx context.Context,mux *runtime.ServeMux,opts []grpc.DialOption) error{
	cfg := initalize.GlobalConfig
	reg := etcd.NewRegistry(registry.Addrs(strings.Split(cfg.Etcd.Address,",")...))
	regSrvs,err := reg.ListServices(func(options *registry.ListOptions) {
		options.Context = ctx
	})
	if err != nil{
		log.Fatal(err)
	}
	//遍历所有etcd注册的服务
	for _,srv := range regSrvs{
		for _,node := range srv.Nodes{
			endpoint := flag.String(srv.Name + uuid.New().String(),node.Address, srv.Name)
			switch srv.Name {
			case constant.API_PROJECT_SRV://项目服务
				err = project.RegisterProjectHandlerFromEndpoint(ctx, mux, *endpoint, opts)

			case constant.API_HELLO_SRV://hello服务
				err = hello.RegisterHelloHandlerFromEndpoint(ctx, mux, *endpoint, opts)
			default:
				return errors.New("srv.Name 服务漏注册了，请解决")
			}
			if err != nil {
				return err
			}
			fmt.Println(srv.Name + " 服务注册端点地址 "+node.Address)
		}
	}
	go func() {
		for{
			time.Sleep(time.Second * 2)
			reg.Watch(func(options *registry.WatchOptions) {
				fmt.Println("11111111")
			})
		}
	}()
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
