package impl

import (
	"go-micro.dev/v4/client"
)

type HelloImpl struct {
	client client.Client
}

func NewHelloImplImpl(cli client.Client)*HelloImpl{
	return &HelloImpl{
		client: cli,
	}
}
