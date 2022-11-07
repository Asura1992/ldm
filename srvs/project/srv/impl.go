package srv

import (
	"go-micro.dev/v4/client"
	"ldm/srvs/project/repos"
)

type ProjectImpl struct {
	client client.Client
	repo   *repos.ProjectModel
}

func NewProjectImpl(cli client.Client, repo *repos.ProjectModel) *ProjectImpl {
	return &ProjectImpl{
		client: cli,
		repo:   repo,
	}
}
