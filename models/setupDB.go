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

	dsn := parseEnv()
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

func parseEnv() (res string) {
	host := os.Getenv("HOST_BD")
	user := os.Getenv("USER_BD")
	password := os.Getenv("PASSWORD_BD")
	database := os.Getenv("DATABASE_BD")
	port := os.Getenv("PORT_BD")
	res = "host=" + host + " user=" + user + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable TimeZone=Europe/Moscow"
	return res
}
