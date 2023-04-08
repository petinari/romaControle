package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendError(ctx *gin.Context, code int, msg interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
