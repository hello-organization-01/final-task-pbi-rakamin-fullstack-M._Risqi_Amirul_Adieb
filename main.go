package main

import (
	"final-task-pbi-rakamin-fullstack-M._Risqi_Amirul_Adieb/helpers"
	"final-task-pbi-rakamin-fullstack-M._Risqi_Amirul_Adieb/router"
	"final-task-pbi-rakamin-fullstack-M._Risqi_Amirul_Adieb/database"
	"github.com/gin-gonic/gin"
)

func init() {
	helpers.LoadEnvVariables()
	database.Connect()
	database.Migration()
}

func main() {
	r := gin.Default()
	router.Routers(r)
	r.Run()
}
