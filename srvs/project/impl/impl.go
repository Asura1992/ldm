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
	fmt.Println("1111111111111111")
	time.Sleep(time.Second * 2)
	rsp.Msg = "project" + req.Name
	fmt.Println("22222222222")
	return nil
}


func NewProjectImpl(cli client.Client)*ProjectImpl{
	return &ProjectImpl{
		client: cli,
	}
}
