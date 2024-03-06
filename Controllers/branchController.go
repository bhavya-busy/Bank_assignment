package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/my/repo/DB"
	models "github.com/my/repo/Models"
)

func CreateBranch(c *gin.Context) {
	var requestBody models.Branch
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := db.DB.Model(&requestBody).Insert(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create branch "})
		return
	}

	c.JSON(http.StatusCreated, requestBody)
}

func ModifyBranch(c *gin.Context) {
	BranchId := c.Param("branch_id")
	var branches *models.Branch
	if err := c.BindJSON(&branches); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch not found"})
		return
	}

	_, err := db.DB.Model(branches).Where("branch_id= ?", BranchId).Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Coudnt Update Branch"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Branch updated successfully"})

}

func DeleteBranch(c *gin.Context) {

	BranchId := c.Param("branch_id")

	_, err := db.DB.Model(&models.Branch{}).Where("branch_id= ?", BranchId).Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Could not delete branch"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Branch deleted successfully"})
}

func GetBranchDetails(c *gin.Context) {
	BranchId := c.Param("branch_id")
	//var banks *models.Bank
	branches := new(models.Branch)
	// if err := c.BindJSON(&banks); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "bank not found"})
	// 	return
	// }

	err := db.DB.Model(branches).Where("branch_id=?", BranchId).Relation("Bank").Relation("Customer").Relation("Account").Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": branches})
}
