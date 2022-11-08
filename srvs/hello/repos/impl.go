package repos

import "gorm.io/gorm"

type HelloModel struct {
	db *gorm.DB
}

func NewHelloModel(db *gorm.DB) *HelloModel {
	return &HelloModel{
		db: db,
	}
}
