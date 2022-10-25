package constant

//允许grpc http请求扩展头部信息,修改后需要重启网关
var MAP_ALLOW_ENTEND_HEADER = map[string]struct{}{
	"Ldm": {},
}
