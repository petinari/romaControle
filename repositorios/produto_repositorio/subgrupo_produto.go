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
		Nome:     subgrupoProduto.Nome,
		IDTenant: uuid.Must(subgrupoProduto.IDTenant, nil),
		IDGrupo:  uuid.Must(subgrupoProduto.IDGrupo, nil),
		Ativo:    subgrupoProduto.Ativo,
	})
	if err != nil {
		return nil, err
	}
	return &subgrupo_inserido, nil
}

// update subgrupo produto
func UpdateSubGrupoProduto(subgrupoProduto db.SubgrupoProduto) (*db.SubgrupoProduto, error) {
	ctx := context.Background()
	queries := db.New(config.DB)
	subgrupoInserido, err := queries.UpdateSubGrupoProduto(ctx, db.UpdateSubGrupoProdutoParams{
		ID:       subgrupoProduto.ID,
		Nome:     subgrupoProduto.Nome,
		IDTenant: subgrupoProduto.IDTenant,
		Ativo:    subgrupoProduto.Ativo,
	})
	if err != nil {
		return nil, err
	}
	return &subgrupoInserido, nil
}

// func get subgrupo produto by grupo produto id
func GetSubGrupoProdutoByGrupoProdutoId(idGrupo uuid.UUID, idTenant uuid.UUID) ([]db.SubgrupoProduto, error) {
	ctx := context.Background()
	queries := db.New(config.DB)
	subgrupoProduto, err := queries.GetSubGrupoProdutoByGrupo(ctx, db.GetSubGrupoProdutoByGrupoParams{
		IDGrupo:  idGrupo,
		IDTenant: idTenant,
	})
	if err != nil {
		return nil, err
	}
	return subgrupoProduto, nil
}
