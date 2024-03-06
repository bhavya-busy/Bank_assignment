package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/my/repo/Controllers"
)

func BranchRoutes(router *gin.Engine) {
	branch := router.Group("branch")
	branch.POST("/", controllers.CreateBranch)
	branch.DELETE("/:branch_id", controllers.DeleteBranch)
	branch.PATCH("/:branch_id", controllers.ModifyBranch)
	branch.GET("/:branch_id", controllers.GetBranchDetails)
}
