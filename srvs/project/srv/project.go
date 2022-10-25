package srv

import (
	"context"
	"ldm/common/protos/project"
)

//获取项目
func (p *ProjectImpl) GetProject(ctx context.Context, req *project.GetProjectReq, rsp *project.GetProjectRsp) error {
	rsp.Msg = "地瓜" + req.Name
	return nil
}
