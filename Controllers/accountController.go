package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/my/repo/DB"
	models "github.com/my/repo/Models"
)

func CreateAccount(c *gin.Context) {
	var requestBody models.Account
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate account details (optional)

	// Create a new bank account record
	// newAccount := models.Account{
	// 	Ac_no: requestBody.Ac_no,
	// 	// Set other account details as needed
	// }

	// Save the new account to the database
	if _, err := db.DB.Model(&requestBody).Insert(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bank account"})
		return
	}

	// Return success response with the created account details
	c.JSON(http.StatusCreated, requestBody)
}

func DeleteAccount(c *gin.Context) {
	AcId := c.Param("ac_id")

	_, err := db.DB.Model(&models.Account{}).Where("ac_id= ?", AcId).Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Could not delete account"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "account deleted successfully"})
}

func ModifyAccount(c *gin.Context) {
	AcId := c.Param("ac_id")
	var accounts *models.Account
	if err := c.BindJSON(&accounts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account not found"})
		return
	}

	_, err := db.DB.Model(accounts).Where("ac_id= ?", AcId).Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Coudnt Update account"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account updated successfully"})
}

func GetAccountDetails(c *gin.Context) {
	AcId := c.Param("ac_id")
	//var banks *models.Bank
	accounts := new(models.Account)
	// if err := c.BindJSON(&banks); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "bank not found"})
	// 	return
	// }

	err := db.DB.Model(accounts).Where("ac_id=?", AcId).Relation("Branch").Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": accounts})
}

