package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"romaControle/servicos/produtos_servicos"
)

func CriarGrupoProduto(ctx *gin.Context) {
	grupoProduto, err := produtos_servicos.CadastrarGrupoProdutos(ctx)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	sendSuccess(ctx, grupoProduto)
}
