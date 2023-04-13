package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"romaControle/controller"
	"romaControle/middleware"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler

	basePath := "/api/"

	base := router.Group(basePath)
	{

		base.POST("/signin", controller.CriarUsuario)
		base.POST("/login", controller.Login)

	}
	cadastrosProduto := fmt.Sprintf("%s%s", basePath, "produtos/")
	produtos := router.Group(cadastrosProduto)
	{
		produtos.POST("/grupos", middleware.RequerAutenticacao, controller.CriarGrupoProduto)
	}

}
