package constant

const (
	//招呼服务
	API_HELLO_SRV = "api-hello-srv"
	//项目服务
	API_PROJECT_SRV = "api-project-srv"
	//直播间服务
	API_LIVEROOM_SRV = "api-liveroom-srv"
)

var MAP_SERVER_ARR = map[string]int{
	API_HELLO_SRV:    9507,
	API_PROJECT_SRV:  9508,
	API_LIVEROOM_SRV: 9509,
}
