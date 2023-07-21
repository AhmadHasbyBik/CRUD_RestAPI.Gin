package main

import (
	"github.com/AhmadHasbyBik/MySQL_RestAPI.Gin/controllers/phoneController"
	"github.com/AhmadHasbyBik/MySQL_RestAPI.Gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnDB()

	r.GET("/api", phoneController.GetAll)
	r.GET("/api/:id", phoneController.GetById)
	r.POST("/api", phoneController.Create)
	r.PUT("/api/:id", phoneController.Update)
	r.DELETE("/api/:id", phoneController.Delete)

	r.Run()
}
