package herocontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustonecrush/api-marvel-database/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var heroes []models.Hero
	
	models.DB.Find(&heroes)
	c.JSON(http.StatusOK, gin.H{"heroes": heroes})

}

func Show(c *gin.Context) {
	
	var hero models.Hero
	id := c.Param("id")

	if err := models.DB.First(&hero, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message" : "Hero is not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"hero": hero})

}

func Create(c *gin.Context) {
	
}

func Update(c *gin.Context) {
	
}

func Delete(c *gin.Context) {
	
}