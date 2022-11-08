package srv

import (
	"go-micro.dev/v4/client"
	"ldm/srvs/hello/repos"
)

type HelloImpl struct {
	client client.Client
	repo   *repos.HelloModel
}

func NewHelloImplImpl(cli client.Client, repo *repos.HelloModel) *HelloImpl {
	return &HelloImpl{
		client: cli,
		repo:   repo,
	}
}
