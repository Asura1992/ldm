package initalize

import (
	"gorm.io/gorm/schema"
	"ldm/common/config"
	"ldm/common/dao"
	"log"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() {
	cfg := config.GlobalConfig.Database
	dsn := cfg.UserName + ":" + cfg.UserPasswd + "@tcp(" + cfg.Address + ")/" + cfg.DbName + "?" + "charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,                         // DSN data source name
		DefaultStringSize:         255,         				// string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: true, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.TablePrefix,   // 表名前缀
			SingularTable: true, // 使用单数表名
		},
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	if cfg.Debug {
		db = db.Debug()
	}
	dao.Db = db
}
