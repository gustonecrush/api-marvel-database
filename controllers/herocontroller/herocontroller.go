package herocontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustonecrush/api-marvel-database/models"
	"gorm.io/gorm"
)

// func Index(c *gin.Context) {

// 	var heroes []models.Hero

// 	models.DB.Find(&heroes)
// 	c.JSON(http.StatusOK, gin.H{"heroes": heroes})

// }

func Index(w http.ResponseWriter, r *http.Request) {

	var heroes []models.Hero
	models.DB.Find(&heroes)

	response, _ := json.Marshal(map[string]any{"heroes": heroes})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

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
	
	var hero models.Hero
	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&hero, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Delete Hero Failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hero": "Data Hero Deleted"})

}