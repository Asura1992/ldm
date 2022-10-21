package impl

import (
	"context"
	"ldm/common/protos/project"
)

//获取项目
func (h ProjectImpl) GetProject(ctx context.Context, req *project.GetProjectReq, rsp *project.GetProjectRsp) error {
	rsp.Msg = "project" + req.Name
	return nil
}
