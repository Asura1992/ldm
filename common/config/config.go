package config
type Config struct {
	//注册中心
	Etcd     		Etcd `mapstructure:"etcd"`
	//对外开放访问端口
	HttpPort 		int  `mapstructure:"http_port"`
	//请求超时
	HttpTimeout 	int  `mapstructure:"http_timeout"`
	//数据库配置
	Database        Database  `mapstructure:"database"`
}
type Etcd struct {
	Address		string `mapstructure:"address"`
}
type Database struct {
	UserName 		string `mapstructure:"user_name"`
	UserPasswd 		string `mapstructure:"user_passwd"`
	Address 		string `mapstructure:"address"`
	DbName 			string `mapstructure:"db_name"`
	Debug           bool   `mapstructure:"debug"`
	TablePrefix     string `mapstructure:"table_prefix"`
}
var GlobalConfig Config
