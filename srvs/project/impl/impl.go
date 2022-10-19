package impl

import (
	"context"
	"fmt"
	"go-micro.dev/v4/client"
	"ldm/common/protos/project"
	"time"
)

type ProjectImpl struct {
	client client.Client
}

func (h ProjectImpl) GetProject(ctx context.Context, req *project.GetProjectReq, rsp *project.GetProjectRsp) error {
	time.Sleep(time.Second * 3)
	rsp.Msg = "project" + req.Name
	fmt.Println("11111111111111")
	return nil
}


func NewProjectImpl(cli client.Client)*ProjectImpl{
	return &ProjectImpl{
		client: cli,
	}
}
