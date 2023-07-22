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
	var phone models.Phone
	if err := c.ShouldBind(&phone); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&phone)
	c.JSON(http.StatusOK, gin.H{"phone": phone})
}

func Update(c *gin.Context) {
	var phone models.Phone
	id := c.Param("id")

	if err := c.ShouldBind(&phone); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&phone).Where("id = ?", id).Updates(&phone).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot Update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Has Been Updated"})
}

func Delete(c *gin.Context) {
	var phone models.Phone
	id := c.Param("id")

	if err := c.ShouldBind(&phone); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Delete(&phone, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot Delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Has Been deleted"})
}
