package srv

import (
	"context"
	"go-micro.dev/v4/registry"
	"ldm/common/constant"
	"log"
)

//监听服务变化
func wathServiceChange(ctx context.Context, reg registry.Registry) error {
	w, err := reg.Watch(func(options *registry.WatchOptions) {
		options.Context = ctx
		//options.Service = constant.API_HELLO_SRV //不写则监听所有服务
	})
	if err != nil {
		return err
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("捕获异常:", err)
			}
		}()
		for {
			rs, err := w.Next()
			if err != nil {
				log.Println("etcd服务监听程序错误:", err)
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
				log.Println(err.Error())
			}
			for _, srv := range srvs {
				if err = registerEndpoint(ctx, *srv); err != nil {
					log.Println("服务:", rs.Service.Name, "重新注册端点失败:", err)
				}
			}
		}
	}()
	return nil
}
