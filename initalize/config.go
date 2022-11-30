package initalize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"ldm/common/config"
	"log"
)

//获取配置
func InitGlobalConfig() {
	configFileName := "config/config.yaml"
	v := viper.New()
	// 设置文件路径
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := v.Unmarshal(&config.GlobalConfig); err != nil {
		log.Fatal(err)
	}
	// 动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Println("配置" + in.Name + "文件产生变化")
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&config.GlobalConfig)
	})
}
