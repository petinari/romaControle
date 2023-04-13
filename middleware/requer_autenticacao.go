package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"romaControle/config"
	"romaControle/seguranca"
)

func RequerAutenticacao(g *gin.Context) {
	er := seguranca.ValidarToken(g)
	if er != nil {
		removeUsuarioDoCache(g)
		g.AbortWithStatus(http.StatusUnauthorized)
		g.Next()
		return
	}
	idU, _, e := seguranca.ExtrairDadosUsuarioToken(g)
	if e != nil {
		removeUsuarioDoCache(g)
		g.AbortWithStatus(http.StatusUnauthorized)
		g.Next()
		return
	}
	redisClient := config.ClientRedis
	val, erro := redisClient.Get(g, idU.String()).Result()
	if erro != nil {
		removeUsuarioDoCache(g)
		g.AbortWithStatus(http.StatusUnauthorized)
		g.Next()
		return

	}
	if val == "" {
		g.AbortWithStatus(http.StatusUnauthorized)
		g.Next()
		return
	}
	g.Next()
}

func removeUsuarioDoCache(g *gin.Context) {

	uId, _, _ := seguranca.ExtrairDadosUsuarioTokenInvalido(g)
	redisClient := config.ClientRedis
	redisClient.GetDel(g, uId.String())
}
