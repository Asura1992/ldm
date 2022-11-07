package srv

import (
	"go-micro.dev/v4/client"
	"ldm/srvs/project/model"
)

type ProjectImpl struct {
	client client.Client
	repo   *model.ProjectModel
}

func NewProjectImpl(cli client.Client, repo *model.ProjectModel) *ProjectImpl {
	return &ProjectImpl{
		client: cli,
		repo:   repo,
	}
}
