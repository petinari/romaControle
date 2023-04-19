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

func GetGrupoProduto(ctx *gin.Context) {
	gruposProdutos, e := produtos_servicos.GetGrupoProdutos(ctx)
	if e != nil {
		sendError(ctx, http.StatusBadRequest, e.Error())
		return
	}
	sendSuccess(ctx, gruposProdutos)
}
