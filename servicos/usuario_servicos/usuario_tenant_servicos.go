package usuario_servicos

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"os"
	"romaControle/config"
	db "romaControle/db/sqlc"
	"romaControle/models/usuario"
	"romaControle/repositorios/usuario_repositorio"
	"romaControle/seguranca"
	"strconv"
	"time"
)

func CriarUsuarioService(ctx *gin.Context) (*usuario.UsuarioOut, error) {

	request := db.Usuario{}
	err := ctx.BindJSON(&request)
	if err != nil {
		return nil, err
	}
	senhaHash, e := seguranca.Hash(request.Senha)
	if e != nil {
		return nil, e
	}
	request.Senha = string(senhaHash)
	usuarioDb, err := usuario_repositorio.CadastrarUsuarioTenant(request)

	if err != nil {
		return nil, err
	}
	usuarioOut := &usuario.UsuarioOut{
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
	return usuarioOut, err
}

func Login(ctx *gin.Context) (interface{}, error) {
	request := usuario.UsuarioLogin{}

	errs := ctx.BindJSON(&request)
	if errs != nil {
		return nil, errs
	}

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
	usuarioOut := &usuario.UsuarioOut{
		ID:    u.ID,
		Email: u.Email,
		TenantOut: usuario.TenantOut{
			IDTenant: u.IDTenant,
		},
	}

	_json, err := json.Marshal(usuarioOut)
	if err != nil {
		return nil, err
	}

	exp, _ := strconv.Atoi(os.Getenv("TEMPO_EXP"))
	err = redisClient.Set(ctx, u.ID.String(), _json, time.Hour*time.Duration(exp)).Err()
	if err != nil {
		return nil, err
	}

	return token, nil
}
