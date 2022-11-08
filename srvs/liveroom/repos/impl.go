package repos

import "gorm.io/gorm"

type LiveroomModel struct {
	db *gorm.DB
}

func NewLiveroomModel(db *gorm.DB) *LiveroomModel {
	return &LiveroomModel{
		db: db,
	}
}
