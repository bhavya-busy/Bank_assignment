package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/my/repo/DB"
	models "github.com/my/repo/Models"
)

func CreateTransaction(c *gin.Context) {
	var requestBody models.Transaction
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if _, err := db.DB.Model(&requestBody).Insert(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, requestBody)
}

func GetTransDetail(c *gin.Context) {
	transId := c.Param("trans_id")
	trans := new(models.Transaction)

	err := db.DB.Model(trans).Where("trans_id=?", transId).Relation("Account").Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": trans})
}

func GetTransByAccount(c *gin.Context) {
	acId := c.Param("ac_id")

	trans := new(models.Transaction)

	err := db.DB.Model(trans).Where("ac_id=?", acId).Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": trans})
}
