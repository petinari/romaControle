package produtos_servicos

import (
	"github.com/gin-gonic/gin"
	db "romaControle/db/sqlc"
	"romaControle/repositorios/produto_repositorio"
	"romaControle/seguranca"
)

// insert subgrupo produto
func CadastrarSubGrupoProdutos(ctx *gin.Context) (*db.SubgrupoProduto, error) {
	request := db.SubgrupoProduto{}
	errs := ctx.BindJSON(&request)
	_, idTenant, err := seguranca.ExtrairDadosUsuarioToken(ctx)
	if err != nil {
		return nil, err
	}
	request.IDTentant = *idTenant
	if errs != nil {
		return nil, errs
	}
	subgrupoInserido, e := produto_repositorio.AddSubGrupoProduto(request)
	if e != nil {
		return nil, e
	}
	return subgrupoInserido, nil
}
