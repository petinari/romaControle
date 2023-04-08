package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"romaControle/servicos"
)

func CriarUsuario(ctx *gin.Context) {
	usuario, err := servicos.CriarUsuarioService(ctx)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	sendSuccess(ctx, usuario)
}
