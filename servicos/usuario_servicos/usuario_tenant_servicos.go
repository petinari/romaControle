package usuario_servicos

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"romaControle/config"
	db "romaControle/db/sqlc"
	"romaControle/models/usuario"
	"romaControle/repositorios/usuario_repositorio"
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
	usuarioDb, err := usuario_repositorio.CadastrarUsuarioTenant(request)

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
	//usuario_repositorio, err := rep.CadastrarUsuario(context.Background(), request)

	//if err != nil {
	//	return nil, err
	//}
	return usuario_out, err
}

func Login(ctx *gin.Context) (interface{}, error) {
	request := usuario.UsuarioLogin{}
	ctx.BindJSON(&request)
	u, e := usuario_repositorio.SelectUsuarioPorEmail(request)
	if e != nil {
		return nil, errors.New("Usuario ou senha incorreto(s)")
	}
	err := seguranca.VerificarSenha(u.Senha, request.Senha)
	if err != nil {
		return nil, errors.New("Usuario ou senha incorreto(s)")
	}
	token, erro := seguranca.CriarToken(u.ID, u.IDTenant)
	if erro != nil {
		return nil, erro
	}
	redisClient := config.ClientRedis
	usuario_out := &usuario.UsuarioOut{
		ID:    u.ID,
		Email: u.Email,
		TenantOut: usuario.TenantOut{
			IDTenant: u.IDTenant,
		},
	}

	json, err := json.Marshal(usuario_out)
	if err != nil {
		return nil, err
	}
	err = redisClient.Set(ctx, u.ID.String(), json, 0).Err()
	if err != nil {
		return nil, err
	}

	return token, nil
}
