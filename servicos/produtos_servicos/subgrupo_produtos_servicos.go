package produtos_servicos

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "romaControle/db/sqlc"
	"romaControle/repositorios/produto_repositorio"
)

// insert subgrupo produto
func CadastrarSubGrupoProdutos(ctx *gin.Context) (*db.SubgrupoProduto, error) {
	request := db.SubgrupoProduto{}
	errs := ctx.BindJSON(&request)
	//_, _, err := seguranca.ExtrairDadosUsuarioToken(ctx)
	//if err != nil {
	//	return nil, err
	//}
	request.IDTenant = uuid.MustParse("9068701a-1386-4b83-9516-5860b2adfeca")
	if errs != nil {
		return nil, errs
	}
	subgrupoInserido, e := produto_repositorio.AddSubGrupoProduto(request)
	if e != nil {
		return nil, e
	}
	return subgrupoInserido, nil
}

// update subgrupo produto
func UpdateSubGrupoProdutos(ctx *gin.Context) (*db.SubgrupoProduto, error) {
	request := db.SubgrupoProduto{}
	errs := ctx.BindJSON(&request)
	//_, _, err := seguranca.ExtrairDadosUsuarioToken(ctx)
	//if err != nil {
	//	return nil, err
	//}
	request.IDTenant = uuid.MustParse("9068701a-1386-4b83-9516-5860b2adfeca")
	if errs != nil {
		return nil, errs
	}
	subgrupoAtualizado, e := produto_repositorio.UpdateSubGrupoProduto(request)
	if e != nil {
		return nil, e
	}
	return subgrupoAtualizado, nil
}

// get subgrupo produto by grupo produto
func GetSubGrupoProdutosByGrupo(ctx *gin.Context) ([]db.SubgrupoProduto, error) {
	idGrupo := ctx.Param("id_grupo")
	//_, _, err := seguranca.ExtrairDadosUsuarioToken(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//request.IDTenant = uuid.MustParse("9068701a-1386-4b83-9516-5860b2adfeca")
	//if errs != nil {
	//	return nil, errs
	//}
	subgrupoProduto, e := produto_repositorio.GetSubGrupoProdutoByGrupoProdutoId(uuid.MustParse(idGrupo), uuid.MustParse("9068701a-1386-4b83-9516-5860b2adfeca"))
	if e != nil {
		return nil, e
	}
	return subgrupoProduto, nil
}
