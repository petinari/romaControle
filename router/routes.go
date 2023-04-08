package router

import (
	"github.com/gin-gonic/gin"
	"romaControle/controller"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler

	basePath := "/api/"

	v1 := router.Group(basePath)
	{

		v1.POST("/usuario", controller.CriarUsuario)

	}

}
