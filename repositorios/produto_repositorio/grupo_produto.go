package produto_repositorio

import (
	"context"
	"github.com/google/uuid"
	"romaControle/config"
	db "romaControle/db/sqlc"
)

func CadastrarGrupoProduto(grupoProduto db.GrupoProduto) (*db.GrupoProduto, error) {
	ctx := context.Background()
	queries := db.New(config.DB)
	grupo_inserido, err := queries.CreateGrupoProdutos(ctx, db.CreateGrupoProdutosParams{
		Nome:     grupoProduto.Nome,
		IDTenant: uuid.Must(grupoProduto.IDTenant, nil),
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
