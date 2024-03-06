package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/my/repo/Controllers"
)

func CustomerRoutes(router *gin.Engine) {
	customer := router.Group("customer")
	customer.POST("/", controllers.CreateCustomer)
	customer.DELETE("/:cust_id", controllers.DeleteCustomer)
	customer.PATCH("/:cust_id", controllers.ModifyCustomer)
	customer.GET("/:cust_id", controllers.GetListOfAccounts)

}
