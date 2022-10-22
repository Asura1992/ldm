package impl

import (
	"context"
	"fmt"
	"ldm/common/protos/project"
)

//获取项目
func (p ProjectImpl) GetProject(ctx context.Context, req *project.GetProjectReq, rsp *project.GetProjectRsp) error {
	rsp.Msg = "project" + req.Name
	fmt.Println("hahahaha")
	return nil
}
