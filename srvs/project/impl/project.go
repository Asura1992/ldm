package impl

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"ldm/common/protos/project"
)

//获取项目
func (p ProjectImpl) GetProject(ctx context.Context, req *project.GetProjectReq, rsp *emptypb.Empty) error {
	return nil
}

