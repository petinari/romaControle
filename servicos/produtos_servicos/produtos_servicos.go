package produtos_servicos

import (
	"github.com/gin-gonic/gin"
	db "romaControle/db/sqlc"
	"romaControle/repositorios/produto_repositorio"
	"romaControle/seguranca"
)

func CadastrarGrupoProdutos(ctx *gin.Context) (*db.GrupoProduto, error) {
	request := db.GrupoProduto{}
	ctx.BindJSON(&request)
	_, idTenant, err := seguranca.ExtrairDadosUsuarioToken(ctx)
	if err != nil {
		return nil, err
	}
	request.IDTenant = *idTenant
	grupo_inserido, e := produto_repositorio.CadastrarGrupoProduto(request)
	if e != nil {
		return nil, e
	}
	return grupo_inserido, nil
}
