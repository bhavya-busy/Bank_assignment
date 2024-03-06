package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/my/repo/DB"
	models "github.com/my/repo/Models"
)

func CreateBank(c *gin.Context) {
	var requestBody models.Bank
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if _, err := db.DB.Model(&requestBody).Insert(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bank "})
		return
	}

	c.JSON(http.StatusCreated, requestBody)
}

func DeleteBank(c *gin.Context) {
	bankId := c.Param("bank_id")

	_, err := db.DB.Model(&models.Bank{}).Where("bank_id= ?", bankId).Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Could not delete bank"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bank deleted successfully"})
}

func ModifyBank(c *gin.Context) {
	bankId := c.Param("bank_id")
	var banks *models.Bank
	if err := c.BindJSON(&banks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bank not found"})
		return
	}

	_, err := db.DB.Model(banks).Where("bank_id= ?", bankId).Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Coudnt Update Bank"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bank updated successfully"})

}

func GetBankDetails(c *gin.Context) {
	bankId := c.Param("bank_id")
	//var banks *models.Bank
	banks := new(models.Bank)
	// if err := c.BindJSON(&banks); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "bank not found"})
	// 	return
	// }

	err := db.DB.Model(banks).Where("bank_id=?", bankId).Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": banks})
}

func GetListOfBranches(c *gin.Context) {
	bankId := c.Param("bank_id")
	//var banks *models.Bank
	banks := new(models.Bank)
	err := db.DB.Model(banks).Where("bank_id=?", bankId).Relation("Branches").Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": banks})
}
