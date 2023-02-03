package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustonecrush/api-marvel-database/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

// func Register(w http.ResponseWriter, r *http.Request) {

// 	// retrieve json data from frontend
// 	var userInput models.User
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&userInput); err != nil {
// 		log.Fatal("Failed to decode json")
// 	}
// 	defer r.Body.Close()

// 	// hash password using bcrypt
// 	hashPassword, _   := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
// 	userInput.Password = string(hashPassword)

// 	// insert to database
// 	if err := models.DB.Create(&userInput).Error; err != nil {
// 		log.Fatal("Failed to register")
// 	}

// 	// create response 
// 	response, _ := json.Marshal(map[string]string{"message": "Success to register"})
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(response)

// }

func Register(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// encrypt the password
	hashPassowrd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password    = string(hashPassowrd)

	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message" : "Success to register"})

}

func Logout(w http.ResponseWriter, r *http.Request) {

}