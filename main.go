package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/gustonecrush/api-marvel-database/controllers/authcontroller"
	"github.com/gustonecrush/api-marvel-database/controllers/herocontroller"
	"github.com/gustonecrush/api-marvel-database/models"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()

	// endpoints hero
	// r.GET("/api/heroes", herocontroller.Index)
	r.GET("/api/heroes/:id", herocontroller.Show)
	r.POST("/api/heroes", herocontroller.Create)
	r.PUT("/api/heroes/:id", herocontroller.Update)
	r.DELETE("/api/heroes", herocontroller.Delete)

	// endpoints auth
	auth := mux.NewRouter()

	auth.HandleFunc("/api/heroes", herocontroller.Index).Methods("GET")
	auth.HandleFunc("/api/login", authcontroller.Login).Methods("POST")
	auth.HandleFunc("/api/register", authcontroller.Register).Methods("POST")
	auth.HandleFunc("/api/logout", authcontroller.Logout).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", auth))

}