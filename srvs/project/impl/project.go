package impl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ldm/common/protos/project"
)

//获取项目
func (h ProjectImpl) GetProject(ctx context.Context, req *project.GetProjectReq, rsp *project.GetProjectRsp) error {
	rsp.Msg = "project" + req.Name
	return status.New(codes.InvalidArgument,"hahahah").Err()
}
