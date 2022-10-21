package impl

import (
	"context"
	"go-micro.dev/v4/client"
	"ldm/common/protos/hello"
)

type HelloImpl struct {
	client client.Client
}

func NewHelloImplImpl(cli client.Client)*HelloImpl{
	return &HelloImpl{
		client: cli,
	}
}
