package authcontroller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gustonecrush/api-marvel-database/config"
	"github.com/gustonecrush/api-marvel-database/helper"
	"github.com/gustonecrush/api-marvel-database/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// take user based on username
	if err := models.DB.Where("username = ?", user.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message" : "Username is not available"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : err.Error()})
			return

		}
	}

	// verify the password inputted with password in database
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message" : "Password is wronng"})
	// 	return
	// }

	// process creating token jwt
	expTime := time.Now().Add(time.Minute * 1)
	claims  := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims:jwt.RegisteredClaims{
			Issuer: "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// declare alogrithm will we use in sign in
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : err.Error()})
		return
	}

	// set token to cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	helper.SendResponse(c, "Successfully Logged In", user)

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

func Logout(c *gin.Context) {

	// remove cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name     : "token",
		Path     : "/",
		Value    : "",
		HttpOnly : true,
		MaxAge   : -1,
	})

	// response
	helper.SendResponse(c, "Successfully logged out", "Logged out success")

}