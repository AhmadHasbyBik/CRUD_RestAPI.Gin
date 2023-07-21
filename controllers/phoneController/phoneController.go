package phoneController

import (
	"net/http"

	"github.com/AhmadHasbyBik/MySQL_RestAPI.Gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(c *gin.Context) {
	var phone []models.Phone

	models.DB.Find(&phone)
	c.JSON(http.StatusOK, gin.H{"phone": phone})
}

func GetById(c *gin.Context) {
	var phone models.Phone
	id := c.Param("id")

	if err := models.DB.First(&phone, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"phone": phone})
}

func Create(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
