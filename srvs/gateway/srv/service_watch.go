package srv

import (
	"context"
	"fmt"
	"go-micro.dev/v4/registry"
	"go.uber.org/zap"
	"ldm/common/constant"
)

//监听服务变化
func wathServiceChange(ctx context.Context, reg registry.Registry) {
	defer func() {
		if err := recover(); err != nil {
			zap.S().Error("捕获异常:",err)
		}
	}()
	w, err := reg.Watch(func(options *registry.WatchOptions) {
		options.Context = ctx
		//options.Service = constant.API_HELLO_SRV //不写则监听所有服务
	})
	if err != nil {
		zap.S().Error("watch service err:",err)
		return
	}
	for {
		rs, err := w.Next()
		if err != nil {
			zap.S().Error("etcd服务监听程序错误:", err.Error())
			return
		}
		//检查是不是自己的服务
		_, ok := constant.MAP_SERVER_ARR[rs.Service.Name]
		if !ok {
			continue
		}
		fmt.Println("检测到服务:", rs.Service.Name, "版本:", rs.Service.Version, "发生变化，动作为:", rs.Action)
		//如果不是创建则跳过
		if rs.Action != "create" {
			continue
		}
		srvs, err := reg.GetService(rs.Service.Name)
		if err != nil {
			zap.S().Error("get service err:",err.Error())
			continue
		}
		for _, srv := range srvs {
			if err = registerEndpoint(ctx, *srv); err != nil {
				zap.S().Error("服务:", rs.Service.Name, "重新注册端点失败:", err.Error())
			}
		}
	}
}
