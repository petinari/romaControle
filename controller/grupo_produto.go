package controller

import (
	"net/http"
	"romaControle/servicos/produtos_servicos"

	"github.com/gin-gonic/gin"
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

// get grupo produto by id and id tenant
func GetGrupoProdutoById(ctx *gin.Context) {
	grupoProduto, er := produtos_servicos.GetGrupoProdutoById(ctx)
	if er != nil {
		sendError(ctx, http.StatusBadRequest, er.Error())
		return
	}
	sendSuccess(ctx, grupoProduto)
}

// disable grupo produto by id and id tenant
func DisableGrupoProdutoById(ctx *gin.Context) {
	grupoProduto, er := produtos_servicos.DisableGrupoProdutoById(ctx)
	if er != nil {
		sendError(ctx, http.StatusBadRequest, er.Error())
		return
	}
	sendSuccess(ctx, grupoProduto)
}
