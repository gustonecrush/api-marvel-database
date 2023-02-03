package main

import (
	"log"
	"net/http"

	// "github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/gustonecrush/api-marvel-database/controllers/authcontroller"
	"github.com/gustonecrush/api-marvel-database/controllers/herocontroller"
	"github.com/gustonecrush/api-marvel-database/models"
)

func main() {

	// r := gin.Default()
	models.ConnectDatabase()

	// endpoints hero
	// r.GET("/api/heroes", herocontroller.Index)
	// r.GET("/api/heroes/:id", herocontroller.Show)
	// r.POST("/api/heroes", herocontroller.Create)
	// r.PUT("/api/heroes/:id", herocontroller.Update)
	// r.DELETE("/api/heroes", herocontroller.Delete)

	// endpoints auth
	r := mux.NewRouter()

	r.HandleFunc("/api/heroes", herocontroller.Index).Methods("GET")
	r.HandleFunc("/api/heroes/{id}", herocontroller.Show).Methods("GET")

	r.HandleFunc("/api/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/api/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/api/logout", authcontroller.Logout).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))

}