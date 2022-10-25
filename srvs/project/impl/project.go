package impl

import (
	"context"
	"fmt"
	"go-micro.dev/v4/metadata"
	"ldm/common/protos/project"
)

//获取项目
func (p ProjectImpl) GetProject(ctx context.Context, req *project.GetProjectReq, rsp *project.GetProjectRsp) error {
	rsp.Msg = "地瓜" + req.Name
	fmt.Println(metadata.FromContext(ctx))
	return nil
}
