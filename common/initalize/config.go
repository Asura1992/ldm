package initalize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)
type config struct {
	Etcd 		Etcd 	`mapstructure:"etcd"`
	HttpPort  	int		`mapstructure:"http_port"`
}
type Etcd struct {
	Address		string `mapstructure:"address"`
}
//全局配置
var GlobalConfig config
//获取配置
func InitGlobalConfig()error{
	configFileName := "config.yaml"
	v := viper.New()
	// 设置文件路径
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		return  err
	}
	if err := v.Unmarshal(&GlobalConfig);err != nil{
		return  err
	}
	// 动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件产生变化：%s", in.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&GlobalConfig)
	})
	return nil
}
