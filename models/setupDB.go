package models

import (
	"fmt"
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
	host := os.Getenv("HOST_DB")
	user := os.Getenv("USER_DB")
	password := os.Getenv("PASSWORD_DB")
	database := os.Getenv("DATABASE_DB")
	port := os.Getenv("PORT_DB")
	res = "host=" + host + " user=" + user + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable TimeZone=Europe/Moscow"
	fmt.Println(res)
	return res
}
