package impl

import (
	"go-micro.dev/v4/client"
	"gorm.io/gorm"
)

type HelloImpl struct {
	client client.Client
	db     *gorm.DB
}

func NewHelloImplImpl(cli client.Client, db *gorm.DB) *HelloImpl {
	return &HelloImpl{
		client: cli,
		db:     db,
	}
}
