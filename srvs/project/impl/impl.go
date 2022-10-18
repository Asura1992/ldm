package impl

import (
	"context"
	"go-micro.dev/v4/client"
	"ldm/common/protos/project"
	"time"
)

type ProjectImpl struct {
	client client.Client
}

func (h ProjectImpl) GetProject(ctx context.Context, req *project.GetProjectReq, rsp *project.GetProjectRsp) error {
	time.Sleep(time.Second * 12)
	rsp.Msg = "project" + req.Name
	return nil
}


func NewProjectImpl(cli client.Client)*ProjectImpl{
	return &ProjectImpl{
		client: cli,
	}
}
