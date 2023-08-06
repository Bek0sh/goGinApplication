package routers

import (
	"project2/controller"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	publicRouter := router.Group("/auth")
	publicRouter.POST("/register", controller.Register)
	publicRouter.POST("/signin", controller.Login)
}
