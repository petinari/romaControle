package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"romaControle/servicos/usuario_servicos"
)

func Login(ctx *gin.Context) {
	token, erro := usuario_servicos.Login(ctx)
	if erro != nil {
		sendError(ctx, http.StatusBadRequest, erro.Error())
		return
	}
	sendSuccess(ctx, token)
}
