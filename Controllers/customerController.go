package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/my/repo/DB"
	models "github.com/my/repo/Models"
)

func CreateCustomer(c *gin.Context) {
	var requestBody models.Customer
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, err := db.DB.Model(&requestBody).Insert(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bank account"})
		return
	}

	c.JSON(http.StatusCreated, requestBody)
}

func DeleteCustomer(c *gin.Context) {
	CustId := c.Param("cust_id")

	_, err := db.DB.Model(&models.Customer{}).Where("cust_id= ?", CustId).Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Could not delete customer"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted successfully"})
}

func ModifyCustomer(c *gin.Context) {
	CustId := c.Param("cust_id")
	var customers *models.Customer
	if err := c.BindJSON(&customers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	_, err := db.DB.Model(customers).Where("cust_id= ?", CustId).Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Coudnt Update customer"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully"})
}

func GetListOfAccounts(c *gin.Context) {
	CustId := c.Param("cust_id")
	//var customers models.Customer
	customers := new(models.Customer)
	err := db.DB.Model(customers).Relation("Map").Where("cust_id=?", CustId).Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Cant fetch "})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": customers})

	////////////

}
