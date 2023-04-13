package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"romaControle/servicos/usuario_servicos"
)

func CriarUsuario(ctx *gin.Context) {
	usuario, err := usuario_servicos.CriarUsuarioService(ctx)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	sendSuccess(ctx, usuario)
}
