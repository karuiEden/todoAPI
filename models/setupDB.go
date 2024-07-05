package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func SetupDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	dsn := "host=localhost user=timeweb password=123 dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, errDB := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if errDB != nil {
		panic("failed to connect database")
	}
	errMig := db.AutoMigrate(&TodoUnit{})
	if errMig != nil {
		panic(errMig)
	}
	DB = db
}
