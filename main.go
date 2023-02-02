package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gustonecrush/api-marvel-database/models"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()

	// routes for API
	r.GET("/api/heroes", herocontroller.Index)
	r.GET("/api/heroes/:id", herocontroller.Show)
	r.POST("/api/heroes", herocontroller.Create)
	r.PUT("/api/heroes/:id", herocontroller.Update)
	r.DELETE("/api/heroes", herocontroller.Delete)

	r.Run()

}