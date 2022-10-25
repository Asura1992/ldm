package srv

import (
	"context"
	"ldm/common/constant"
	"ldm/common/protos/hello"
	"ldm/common/protos/project"
)

//获取项目
func (p *ProjectImpl) GetProject(ctx context.Context, req *project.GetProjectReq, rsp *project.GetProjectRsp) error {
	r, _ := hello.NewHelloService(constant.API_HELLO_SRV, p.client).Hello(ctx, &hello.HelloReq{Name: "45748785488"})
	rsp.Msg = r.Msg
	return nil
}
