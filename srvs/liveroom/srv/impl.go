package srv

import (
	"go-micro.dev/v4/client"
	"ldm/srvs/liveroom/repos"
)

type LiveroomImpl struct {
	client client.Client
	repo   *repos.LiveroomModel
}

func NewLiveroomImpl(cli client.Client, repo *repos.LiveroomModel) *LiveroomImpl {
	return &LiveroomImpl{
		client: cli,
		repo:   repo,
	}
}
