package config

var GlobalConfig Config

type Config struct {
	//版本
	Version string `json:"version"`
	//对外开放访问端口
	HttpPort int `mapstructure:"http_port"`
	//请求超时
	HttpTimeout int `mapstructure:"http_timeout"`
	//注册中心
	Etcd Etcd `mapstructure:"etcd"`
	//mysql数据库配置
	Database Database `mapstructure:"database"`
	//redis
	Redis Redis `mapstructure:"redis"`
	//mns
	Mns Mns `mapstructure:"mns"`
	//链路追踪
	Jaeger Jaeger `mapstructure:"jaeger"`
	//swagger
	Swagger Swagger `mapstructure:"swagger"`
}
type Swagger struct {
	SwaggerAddr string `mapstructure:"swagger_addr"`
}

type Jaeger struct {
	JaegerTracerAddr string `mapstructure:"jaeger_tracer_addr"`
}

type Etcd struct {
	Address string `mapstructure:"address"`
}
type Database struct {
	UserName    string `mapstructure:"user_name"`
	UserPasswd  string `mapstructure:"user_passwd"`
	Address     string `mapstructure:"address"`
	DbName      string `mapstructure:"db_name"`
	Debug       bool   `mapstructure:"debug"`
	TablePrefix string `mapstructure:"table_prefix"`
}
type Redis struct {
	Address string `mapstructure:"address"`
	Db      int    `mapstructure:"db"`
}
type Mns struct {
	Url             string `mapstructure:"url"`
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	Env             string `mapstructure:"env"`
}
