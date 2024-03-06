package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/my/repo/Controllers"
)

func BankRoutes(router *gin.Engine) {
	bank := router.Group("bank")
	bank.POST("/", controllers.CreateBank)
	bank.DELETE("/:bank_id", controllers.DeleteBank)
	bank.PATCH("/:bank_id", controllers.ModifyBank)
	bank.GET("/:bank_id", controllers.GetBankDetails)
	bank.GET("/list/:bank_id", controllers.GetListOfBranches)
}
