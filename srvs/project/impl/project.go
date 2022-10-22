package impl

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"ldm/common/protos/project"
)

//获取项目
func (p ProjectImpl) GetProject(ctx context.Context, req *project.GetProjectReq, rsp *project.GetProjectRsp) error {
	rsp.Msg =  "地瓜" + req.Name
	m,_ := metadata.FromIncomingContext(ctx)
	fmt.Println(m)
	return nil
}
