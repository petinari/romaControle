package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"romaControle/servicos"
)

func Login(ctx *gin.Context) {
	token, erro := servicos.Login(ctx)
	if erro != nil {
		sendError(ctx, http.StatusBadRequest, erro.Error())
		return
	}
	sendSuccess(ctx, token)
}
