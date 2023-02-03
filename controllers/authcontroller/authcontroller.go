package authcontroller

import (
	"encoding/json"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/gustonecrush/api-marvel-database/models"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {

	// retrieve json data from frontend
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		log.Fatal("Failed to decode json")
	}
	defer r.Body.Close()

	// hash password using bcrypt
	hashPassword, _   := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert to database
	if err := models.DB.Create(&userInput).Error; err != nil {
		log.Fatal("Failed to register")
	}

	response, _ := json.Marshal(map[string]string{"message": "Success to register"})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func Logout(w http.ResponseWriter, r *http.Request) {

}