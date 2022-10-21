package impl

import (
	"context"
	"errors"
	"ldm/common/protos/project"
)

//获取项目
func (h ProjectImpl) GetProject(ctx context.Context, req *project.GetProjectReq, rsp *project.GetProjectRsp) error {
	rsp.Msg = "project" + req.Name
	return errors.New("hahahah")
}
