package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/my/repo/Controllers"
)

func TransactionRoutes(router *gin.Engine) {
	trans := router.Group("transaction")
	trans.POST("/", controllers.CreateTransaction)
	trans.GET("/:trans_id", controllers.GetTransDetail)
	trans.GET("/acc/:ac_id", controllers.GetTransByAccount)

}
