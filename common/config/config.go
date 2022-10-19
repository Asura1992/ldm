package config
type Config struct {
	Etcd     		Etcd `mapstructure:"etcd"`
	HttpPort 		int  `mapstructure:"http_port"`
	HttpTimeout 	int  `mapstructure:"http_timeout"`
}
type Etcd struct {
	Address		string `mapstructure:"address"`
}
var GlobalConfig Config
