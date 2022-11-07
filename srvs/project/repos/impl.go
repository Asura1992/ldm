package repos

import (
	"gorm.io/gorm"
)

type ProjectModel struct {
	db *gorm.DB
}

func NewProjectModel(db *gorm.DB) *ProjectModel {
	return &ProjectModel{
		db: db,
	}
}
