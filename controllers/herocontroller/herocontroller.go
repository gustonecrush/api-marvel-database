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

	var hero models.Hero

	if err := c.ShouldBindJSON(&hero); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&hero)
	c.JSON(http.StatusOK, gin.H{"hero": hero})
	
}

func Update(c *gin.Context) {

	var hero models.Hero
	id := c.Param("id")

	if err := c.ShouldBindJSON(&hero); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&hero).Where("id = ?", id).Updates(&hero).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Update Hero Failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hero": "Data Hero Updated"})
	
}

func Delete(c *gin.Context) {
	
}