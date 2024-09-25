package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func InitDbConnection() {
	level := AppConfig.Log.GetLevel()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      level,
			Colorful:      true,
		},
	)

	path := AppConfig.Db.Path
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("sql路径错误!")
	}
	Db = db
}
