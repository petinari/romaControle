package servicos

import (
	"context"
	"github.com/gin-gonic/gin"
	"romaControle/config"
	"romaControle/repositorios"
	"romaControle/schemas"
	"romaControle/seguranca"
)

func CriarUsuarioService(ctx *gin.Context) (*schemas.Usuario, error) {
	request := schemas.Usuario{}
	ctx.BindJSON(&request)

	senhaHash, e := seguranca.Hash(request.Senha)
	if e != nil {
		return nil, e
	}
	request.Senha = string(senhaHash)
	rep := repositorios.NovoRepUser(config.Client.Database(ctx.Request.Header.Get("tenant")))
	usuario, err := rep.CadastrarUsuario(context.Background(), request)

	if err != nil {
		return nil, err
	}
	return usuario, err
}
