package srv

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"ldm/common/constant"
	"ldm/common/protos/hello"
	"ldm/common/protos/project"
)

//获取项目
func (p *ProjectImpl) GetProject(ctx context.Context, req *project.GetProjectReq, rsp *project.GetProjectRsp) error {
	r, _ := hello.NewHelloService(constant.API_HELLO_SRV, p.client).Hello(ctx, &hello.HelloReq{Name: req.Name})
	rsp.Msg = r.Msg
	return nil
}

func (p *ProjectImpl) DelProject(ctx context.Context, req *emptypb.Empty, rsp *emptypb.Empty) error {
	return nil
}
