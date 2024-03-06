package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/my/repo/Controllers"
)

func AccountRoutes(router *gin.Engine) {
	account := router.Group("account")
	account.POST("/", controllers.CreateAccount)
	account.DELETE("/:ac_id", controllers.DeleteAccount)
	account.PATCH("/:ac_id", controllers.ModifyAccount)
	account.GET("/:ac_id", controllers.GetAccountDetails)
}
