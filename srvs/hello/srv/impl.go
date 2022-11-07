package srv

import (
	"go-micro.dev/v4/client"
	"ldm/srvs/hello/model"
)

type HelloImpl struct {
	client client.Client
	repo   *model.HelloModel
}

func NewHelloImplImpl(cli client.Client, repo *model.HelloModel) *HelloImpl {
	return &HelloImpl{
		client: cli,
		repo:   repo,
	}
}
