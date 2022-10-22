package impl

import (
	"go-micro.dev/v4/client"
)

type ProjectImpl struct {
	client client.Client
}

func NewProjectImpl(cli client.Client)*ProjectImpl{
	return &ProjectImpl{
		client: cli,
	}
}
