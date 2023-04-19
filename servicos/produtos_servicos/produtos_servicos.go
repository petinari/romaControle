package produtos_servicos

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "romaControle/db/sqlc"
	"romaControle/repositorios/produto_repositorio"
	"romaControle/seguranca"
)

func CadastrarGrupoProdutos(ctx *gin.Context) (*db.GrupoProduto, error) {
	request := db.GrupoProduto{}
	errs := ctx.BindJSON(&request)
	if errs != nil {
		return nil, errs
	}
	_, idTenant, err := seguranca.ExtrairDadosUsuarioToken(ctx)
	if err != nil {
		return nil, err
	}
	request.IDTenant = *idTenant
	grupoInserido, e := produto_repositorio.CadastrarGrupoProduto(request)
	if e != nil {
		return nil, e
	}
	return grupoInserido, nil
}

func GetGrupoProdutos(ctx *gin.Context) (*[]db.GrupoProduto, error) {
	gruposProdutos, e := produto_repositorio.GetGrupoProduto(uuid.MustParse("876f9634-0ffe-4e60-a66b-a1e6c1aec120"))
	if e != nil {
		return nil, e
	}
	return &gruposProdutos, nil
}
