package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/my/repo/DB"
	models "github.com/my/repo/Models"
)

func CreateMap(c *gin.Context) {
	var requestBody models.Map
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := db.DB.Model(&requestBody).Insert(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create map"})
		return
	}

	c.JSON(http.StatusCreated, requestBody)

}

func DeleteMap(c *gin.Context) {
	mId := c.Param("map_id")

	_, err := db.DB.Model(&models.Map{}).Where("map_id= ?", mId).Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Could not delete "})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
