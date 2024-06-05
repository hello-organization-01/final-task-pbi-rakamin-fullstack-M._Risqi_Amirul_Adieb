package database

import (
	"final-task-pbi-rakamin-fullstack-M._Risqi_Amirul_Adieb/models"
)

func Migration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Photo{})
}
