package impl

import (
	"context"
	"go-micro.dev/v4/client"
	"ldm/common/protos/hello"
)

type HelloImpl struct {
	client client.Client
}

func (h HelloImpl) Hello(ctx context.Context, req *hello.HelloReq, rsp *hello.HelloRsp) error {
	rsp.Msg = "hello" + req.Name
	return nil
}

func NewHelloImplImpl(cli client.Client)*HelloImpl{
	return &HelloImpl{
		client: cli,
	}
}
