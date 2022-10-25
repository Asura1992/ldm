package dao

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Redis *redis.Client
