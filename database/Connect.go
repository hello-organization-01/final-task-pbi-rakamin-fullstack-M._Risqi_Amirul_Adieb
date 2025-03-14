package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(mysql.Open(os.Getenv("DB")), &gorm.Config{})

	if err != nil {
		panic(" Could not connect mysql DB ")
	}
}
