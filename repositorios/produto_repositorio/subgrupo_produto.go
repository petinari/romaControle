package produto_repositorio

import (
	"context"
	"github.com/google/uuid"
	"romaControle/config"
	db "romaControle/db/sqlc"
)

// add subgrupo produto
func AddSubGrupoProduto(subgrupoProduto db.SubgrupoProduto) (*db.SubgrupoProduto, error) {
	ctx := context.Background()
	queries := db.New(config.DB)
	subgrupo_inserido, err := queries.InsertSubGrupoProduto(ctx, db.InsertSubGrupoProdutoParams{
		Nome:      subgrupoProduto.Nome,
		IDTentant: uuid.Must(subgrupoProduto.IDTentant, nil),
		IDGrupo:   uuid.Must(subgrupoProduto.IDGrupo, nil),
		Ativo:     subgrupoProduto.Ativo,
	})
	if err != nil {
		return nil, err
	}
	return &subgrupo_inserido, nil
}
