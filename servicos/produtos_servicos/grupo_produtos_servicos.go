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
	//_, _, err := seguranca.ExtrairDadosUsuarioToken(ctx)
	//if err != nil {
	//	return nil, err
	//}
	request.IDTenant = uuid.MustParse("9068701a-1386-4b83-9516-5860b2adfeca")
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
	//_, _, err := seguranca.ExtrairDadosUsuarioToken(ctx)
	//if err != nil {
	//	return nil, err
	//}
	grupoProduto, er := produto_repositorio.GetGrupoProdutoById(uuid.MustParse(idGrupo), uuid.MustParse("9068701a-1386-4b83-9516-5860b2adfeca"))
	if er != nil {
		return nil, er
	}
	return grupoProduto, nil
}

// disable grupo produto by id and id tenant
func DisableGrupoProdutoById(ctx *gin.Context) (*string, error) {
	idGrupo := ctx.Param("id_grupo")
	//_, _, err := seguranca.ExtrairDadosUsuarioToken(ctx)
	//if err != nil {
	//	return nil, err
	//}
	_, er := produto_repositorio.DisableGrupoProdutoById(uuid.MustParse(idGrupo), uuid.MustParse("9068701a-1386-4b83-9516-5860b2adfeca"))
	if er != nil {
		return nil, er
	}
	msg := "Grupo de produto desabilitado com sucesso"
	return &msg, nil
}

// get grupo produto by nome and id tenant
func GetGrupoProdutoByNome(ctx *gin.Context) (*db.SelectGrupoProdutosByNomeRow, error) {
	nome := ctx.Param("nome")
	_, _, err := seguranca.ExtrairDadosUsuarioToken(ctx)
	if err != nil {
		return nil, err
	}
	grupoProduto, er := produto_repositorio.GetGrupoProdutoByNome(nome, uuid.MustParse("9068701a-1386-4b83-9516-5860b2adfeca"))
	if er != nil {
		return nil, er
	}
	return grupoProduto, nil
}

// update grupo produto by id and id tenant
func UpdateGrupoProdutoById(ctx *gin.Context) (*db.GrupoProduto, error) {
	//idGrupo := ctx.Param("id_grupo")
	//_, _, err := seguranca.ExtrairDadosUsuarioToken(ctx)
	//if err != nil {
	//	return nil, err
	//}
	request := db.GrupoProduto{}
	request.IDTenant = uuid.MustParse("9068701a-1386-4b83-9516-5860b2adfeca")
	errs := ctx.BindJSON(&request)
	if errs != nil {
		return nil, errs
	}
	grupoProduto, er := produto_repositorio.UpdateGrupoProdutoById(request)
	if er != nil {
		return nil, er
	}
	return grupoProduto, nil
}
