package main

import (
	"context"
	"encoding/json"
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
	"ldm/common/proto/gateway/protos/hello"
	"ldm/common/proto/gateway/protos/project"
	"ldm/initalize"
	"ldm/utils/grpc_err"
	"log"
	"net/http"
	"strings"
	"time"
)
var mux = runtime.NewServeMux(
	//允许所有头信息
	runtime.WithIncomingHeaderMatcher(allowHeader),
	//错误响应组装器
	runtime.WithErrorHandler(errResponseBuilder),
	runtime.WithMarshalerOption(
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

//失败请求响应组装器
func errResponseBuilder(ctx context.Context, serveMux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
	errMsg := strings.ReplaceAll(err.Error(),"rpc error: code = Unknown desc = ","")
	if errMsg == ""{
		b,_ := json.Marshal(grpc_err.SelfDefineErr{
			Code: -1,
			Message: "unknown err",
		})
		writer.Write(b)
		return
	}
	var errInfo grpc_err.SelfDefineErr
	if err = json.Unmarshal([]byte(errMsg),&errInfo);err != nil{
		b,_ := json.Marshal(grpc_err.SelfDefineErr{
			Code: -1,
			Message: errMsg,
		})
		writer.Write(b)
		return
	}
	if errInfo.Message == ""{
		errInfo.Message = errMsg
		errInfo.Code = -1
	}
	if errInfo.Code == 0{
		errInfo.Code = -1
	}
	f,_ := json.Marshal(errInfo)
	writer.Write(f)
}
//允许哪些自定义头信息
func allowHeader(s string) (string, bool) {
	switch s {
	case "Ldm":
		return s,true
	}
	return "",false
}
//初始化网关
func initGateway() error{
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	cfg := config.GlobalConfig
	reg := etcd.NewRegistry(registry.Addrs(strings.Split(cfg.Etcd.Address,",")...))
	regSrvs,err := reg.ListServices(func(options *registry.ListOptions) {
		options.Context = ctx
	})
	if err != nil{
		log.Fatal(err)
	}
	//遍历所有etcd注册的服务
	for _,srv := range regSrvs{
		if err = registerEndpoint(ctx,*srv);err != nil{
			return err
		}
	}
	//监听服务变化重新注册端点
	wathServiceChange(ctx,reg)
	//http监听服务启动
	listenAddr := fmt.Sprintf(":%d",config.GlobalConfig.HttpPort)
	connectTimeout := time.Second * time.Duration(config.GlobalConfig.HttpTimeout)
	return http.ListenAndServe(listenAddr, http.TimeoutHandler(mux,connectTimeout,"request timeout o(╥﹏╥)o"))
}
//注册端点
func registerEndpoint(ctx context.Context,srv registry.Service)(err error){
	opts := []grpc.DialOption{grpc.WithInsecure()}
	for _,node := range srv.Nodes{
		endpoint := flag.String(srv.Name + uuid.New().String(),node.Address, srv.Name)
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
	return nil
}
//监听服务变化
func wathServiceChange(ctx context.Context,reg registry.Registry) error{
	w,err := reg.Watch(func(options *registry.WatchOptions) {
		options.Context = ctx
		//options.Service = constant.API_HELLO_SRV //不写则监听所有服务
	})
	if err != nil{
		return err
	}
	go func() {
		defer func() {
			if err := recover();err != nil{
				log.Println("捕获异常:",err)
			}
		}()
		for {
			rs ,err := w.Next()
			if err != nil{
				log.Println("etcd服务监听程序错误:",err)
				return
			}
			fmt.Println(rs.Service.Name,"服务发生变化，变化动作为:",rs.Action)
			//如果不是创建则跳过
			if rs.Action != "create"{
				continue
			}
			srvs,err := reg.GetService(rs.Service.Name)
			if err != nil{
				log.Println(err.Error())
			}
			for _,srv := range srvs{
				if err = registerEndpoint(ctx,*srv);err  != nil{
					log.Println("服务:",rs.Service.Name,"重新注册端点失败:",err)
				}
			}
		}
	}()
	return nil
}
func main() {
	//初始化配置
	initalize.InitGlobalConfig()
	if err := initGateway(); err != nil {
		log.Fatal(err)
	}
}
