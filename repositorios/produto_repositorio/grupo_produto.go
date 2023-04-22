package produto_repositorio

import (
	"context"
	"romaControle/config"
	db "romaControle/db/sqlc"

	"github.com/google/uuid"
)

func CadastrarGrupoProduto(grupoProduto db.GrupoProduto) (*db.GrupoProduto, error) {
	ctx := context.Background()
	queries := db.New(config.DB)
	grupo_inserido, err := queries.CreateGrupoProdutos(ctx, db.CreateGrupoProdutosParams{
		Nome:     grupoProduto.Nome,
		IDTenant: uuid.Must(grupoProduto.IDTenant, nil),
		Ativo:    grupoProduto.Ativo,
	})
	if err != nil {
		return nil, err
	}
	return &grupo_inserido, nil
}

func GetGrupoProduto(id_tenant uuid.UUID) ([]db.GrupoProduto, error) {
	ctx := context.Background()
	queries := db.New(config.DB)
	gruposProdutos, er := queries.SelectGrupoProdutos(ctx, id_tenant)
	if er != nil {
		return nil, er
	}
	return gruposProdutos, nil
}

// get grupo produto by id and id tenant
func GetGrupoProdutoById(id_grupo uuid.UUID, id_tenant uuid.UUID) (*db.GrupoProduto, error) {
	ctx := context.Background()
	queries := db.New(config.DB)
	grupoProduto, er := queries.SelectGrupoProdutosById(ctx, db.SelectGrupoProdutosByIdParams{
		ID:       id_grupo,
		IDTenant: id_tenant,
	})
	if er != nil {
		return nil, er
	}
	return &grupoProduto, nil
}

// disable grupo produto by id and id tenant
func DisableGrupoProdutoById(id_grupo uuid.UUID, id_tenant uuid.UUID) (*db.GrupoProduto, error) {
	ctx := context.Background()
	queries := db.New(config.DB)
	grupoProduto, er := queries.DisableGrupoProdutos(ctx, db.DisableGrupoProdutosParams{
		ID:       id_grupo,
		IDTenant: id_tenant,
	})
	if er != nil {
		return nil, er
	}
	return &grupoProduto, nil
}
