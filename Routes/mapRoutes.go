// package routes
// import (
// 	"github.com/gin-gonic/gin"
// 	controllers "github.com/my/repo/Controllers"
// )

// func Maproutes(router *gin.Engine) {
// 	map:= router.Group("map")
// 	map.POST("/", controllers.CreateMap)
// 	map.DELETE("/:map_id", controllers.DeleteMap)

// }

package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/my/repo/Controllers"
)

func MapRoutes(router *gin.Engine) {
	mapR := router.Group("map")
	mapR.POST("/", controllers.CreateMap)
	mapR.DELETE("/:map_id", controllers.DeleteMap)

}
