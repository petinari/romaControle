package produtos_servicos

import (
	db "romaControle/db/sqlc"
	"romaControle/repositorios/produto_repositorio"
	"romaControle/seguranca"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	gruposProdutos, e := produto_repositorio.GetGrupoProduto(uuid.MustParse("9068701a-1386-4b83-9516-5860b2adfeca"))
	if e != nil {
		return nil, e
	}
	return &gruposProdutos, nil
}

// get grupo produto by id and id tenant
func GetGrupoProdutoById(ctx *gin.Context) (*db.GrupoProduto, error) {
	idGrupo := ctx.Param("id_grupo")
	_, tenantID, err := seguranca.ExtrairDadosUsuarioToken(ctx)
	if err != nil {
		return nil, err
	}
	grupoProduto, er := produto_repositorio.GetGrupoProdutoById(uuid.MustParse(idGrupo), *tenantID)
	if er != nil {
		return nil, er
	}
	return grupoProduto, nil
}

// disable grupo produto by id and id tenant
func DisableGrupoProdutoById(ctx *gin.Context) (*db.GrupoProduto, error) {
	idGrupo := ctx.Param("id_grupo")
	_, tenantID, err := seguranca.ExtrairDadosUsuarioToken(ctx)
	if err != nil {
		return nil, err
	}
	grupoProduto, er := produto_repositorio.DisableGrupoProdutoById(uuid.MustParse(idGrupo), *tenantID)
	if er != nil {
		return nil, er
	}
	return grupoProduto, nil
}
