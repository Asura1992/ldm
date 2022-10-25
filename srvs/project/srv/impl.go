package srv

import (
	"go-micro.dev/v4/client"
	"gorm.io/gorm"
)

type ProjectImpl struct {
	client client.Client
	db     *gorm.DB
}

func NewProjectImpl(cli client.Client, db *gorm.DB) *ProjectImpl {
	return &ProjectImpl{
		client: cli,
		db:     db,
	}
}
