package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"romaControle/servicos/produtos_servicos"
)

// create a new subgrupo produto
func CriarSubGrupoProduto(ctx *gin.Context) {
	subgrupoProduto, err := produtos_servicos.CadastrarSubGrupoProdutos(ctx)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	sendSuccess(ctx, subgrupoProduto)
}

// update subgrupo produto
func UpdateSubGrupoProduto(ctx *gin.Context) {
	subgrupoProduto, err := produtos_servicos.UpdateSubGrupoProdutos(ctx)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	sendSuccess(ctx, subgrupoProduto)
}

// get subgrupo produto by grupo produto
func GetSubGrupoProdutoByGrupo(ctx *gin.Context) {
	subgrupoProduto, err := produtos_servicos.GetSubGrupoProdutosByGrupo(ctx)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	sendSuccess(ctx, subgrupoProduto)
}
