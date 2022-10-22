package impl

import (
	"context"
	"fmt"
	"go-micro.dev/v4/metadata"
	"ldm/common/constant"
	"ldm/common/protos/hello"
	"ldm/common/protos/project"
)

//获取项目
func (p ProjectImpl) GetProject(ctx context.Context, req *project.GetProjectReq, rsp *project.GetProjectRsp) error {
	fmt.Println(metadata.FromContext(ctx))
	r,_ := hello.NewHelloService(constant.API_HELLO_SRV,p.client).Hello(ctx,&hello.HelloReq{Name: "hjsdjfhksjfh"})
	rsp.Msg = r.Msg
	return nil
}
