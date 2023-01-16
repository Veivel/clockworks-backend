package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type dbCredentials struct {
	username string
	password string
	host     string
	port     string
	dbName   string
}

func ConnectDB() {
	var creds = dbCredentials{
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		dbName:   os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		creds.username,
		creds.password,
		creds.host,
		creds.port,
		creds.dbName,
	)

	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	DB = database

	DB.AutoMigrate(&Event{}) // 'events' table
	DB.AutoMigrate(&User{})  // 'users' table
	DB.AutoMigrate(&Tag{})   // 'tags' table

	fmt.Println("Succesfully connected to database.")
}
