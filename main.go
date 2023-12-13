package main

import (
	"golang_talent/controllers/talentcontroller"
	"golang_talent/models"
	"golang_talent/services"

	"github.com/gin-gonic/gin"
)

// * initial .env
func init() {
	services.InitEnv()
}

func main() {

	models.ConnectDatabase()
	r := gin.Default()

	r.GET("/api-talent/talent", talentcontroller.Index)
	r.GET("/api-talent/talent/:id", talentcontroller.Show)
	r.POST("/api-talent/create", talentcontroller.Create)
	r.Run()
}
