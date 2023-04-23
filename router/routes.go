package router

import (
	"fmt"
	"romaControle/controller"

	"github.com/gin-gonic/gin"
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
		produtos.POST("/grupos", controller.CriarGrupoProduto)
		produtos.GET("/grupos/:id_grupo", controller.GetGrupoProdutoById)
		produtos.GET("/grupos", controller.GetGrupoProduto)
		produtos.PUT("/grupos/:id_grupo", controller.DisableGrupoProdutoById)
		produtos.GET("/grupos/nome/:nome", controller.GetGrupoProdutoByNome)
		produtos.POST("/subgrupos", controller.CriarSubGrupoProduto)

	}

}
