package impl

import (
	"context"
	"ldm/common/protos/project"
	"net/http"
)

//获取项目
func (p ProjectImpl) GetProject(ctx context.Context, req *project.GetProjectReq, rsp *project.GetProjectRsp) error {
	rsp.Msg = "project" + req.Name
	return CommonGrpcErr(http.StatusUnauthorized,"hahahah")
}

