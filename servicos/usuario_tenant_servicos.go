package servicos

import (
	"errors"
	"github.com/gin-gonic/gin"
	db "romaControle/db/sqlc"
	"romaControle/models/usuario"
	usuarioRepo "romaControle/repositorios/usuario"
	"romaControle/seguranca"
)

func CriarUsuarioService(ctx *gin.Context) (*usuario.UsuarioOut, error) {

	request := db.Usuario{}
	ctx.BindJSON(&request)
	senhaHash, e := seguranca.Hash(request.Senha)
	if e != nil {
		return nil, e
	}
	request.Senha = string(senhaHash)
	usuarioDb, err := usuarioRepo.CadastrarUsuarioTenant(request)

	if err != nil {
		return nil, err
	}
	usuario_out := &usuario.UsuarioOut{
		ID:    usuarioDb.ID,
		Email: usuarioDb.Email,
		TenantOut: usuario.TenantOut{
			IDTenant: usuarioDb.IDTenant,
		},
	}

	//rep := repositorios.NovoRepUser(config.Client.Database(ctx.Request.Header.Get("tenant")))
	//usuario, err := rep.CadastrarUsuario(context.Background(), request)

	//if err != nil {
	//	return nil, err
	//}
	return usuario_out, err
}

func Login(ctx *gin.Context) (interface{}, error) {
	request := usuario.UsuarioLogin{}
	ctx.BindJSON(&request)
	u, e := usuarioRepo.SelectUsuarioPorEmail(request)
	if e != nil {
		return nil, errors.New("Usuario ou senha incorreto(s)")
	}
	err := seguranca.VerificarSenha(request.Senha, u.Senha)
	if err != nil {
		return nil, errors.New("Usuario ou senha incorreto(s)")
	}
	return u, nil
}
