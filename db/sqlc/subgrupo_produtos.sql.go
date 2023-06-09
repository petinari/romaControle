// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: subgrupo_produtos.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const ativarSubGrupoProduto = `-- name: AtivarSubGrupoProduto :one
update public.subgrupo_produtos set ativo = true where id = $1 and id_tenant = $2 returning id, nome, id_grupo, ativo, id_tenant
`

type AtivarSubGrupoProdutoParams struct {
	ID       uuid.UUID `json:"id"`
	IDTenant uuid.UUID `json:"id_tenant"`
}

func (q *Queries) AtivarSubGrupoProduto(ctx context.Context, arg AtivarSubGrupoProdutoParams) (SubgrupoProduto, error) {
	row := q.db.QueryRow(ctx, ativarSubGrupoProduto, arg.ID, arg.IDTenant)
	var i SubgrupoProduto
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.IDGrupo,
		&i.Ativo,
		&i.IDTenant,
	)
	return i, err
}

const desativarSubGrupoProduto = `-- name: DesativarSubGrupoProduto :one
update public.subgrupo_produtos set ativo = false where id = $1 and id_tenant = $2 returning id, nome, id_grupo, ativo, id_tenant
`

type DesativarSubGrupoProdutoParams struct {
	ID       uuid.UUID `json:"id"`
	IDTenant uuid.UUID `json:"id_tenant"`
}

func (q *Queries) DesativarSubGrupoProduto(ctx context.Context, arg DesativarSubGrupoProdutoParams) (SubgrupoProduto, error) {
	row := q.db.QueryRow(ctx, desativarSubGrupoProduto, arg.ID, arg.IDTenant)
	var i SubgrupoProduto
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.IDGrupo,
		&i.Ativo,
		&i.IDTenant,
	)
	return i, err
}

const getSubGrupoProdutoByGrupo = `-- name: GetSubGrupoProdutoByGrupo :many
select id, nome, id_grupo, ativo, id_tenant from public.subgrupo_produtos where id_grupo = $1 and id_tenant = $2 and ativo = true
`

type GetSubGrupoProdutoByGrupoParams struct {
	IDGrupo  uuid.UUID `json:"id_grupo"`
	IDTenant uuid.UUID `json:"id_tenant"`
}

func (q *Queries) GetSubGrupoProdutoByGrupo(ctx context.Context, arg GetSubGrupoProdutoByGrupoParams) ([]SubgrupoProduto, error) {
	rows, err := q.db.Query(ctx, getSubGrupoProdutoByGrupo, arg.IDGrupo, arg.IDTenant)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SubgrupoProduto{}
	for rows.Next() {
		var i SubgrupoProduto
		if err := rows.Scan(
			&i.ID,
			&i.Nome,
			&i.IDGrupo,
			&i.Ativo,
			&i.IDTenant,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSubGrupoProdutoById = `-- name: GetSubGrupoProdutoById :one
select id, nome, id_grupo, ativo, id_tenant from public.subgrupo_produtos where id = $1 and id_tenant = $2 and ativo = true
`

type GetSubGrupoProdutoByIdParams struct {
	ID       uuid.UUID `json:"id"`
	IDTenant uuid.UUID `json:"id_tenant"`
}

func (q *Queries) GetSubGrupoProdutoById(ctx context.Context, arg GetSubGrupoProdutoByIdParams) (SubgrupoProduto, error) {
	row := q.db.QueryRow(ctx, getSubGrupoProdutoById, arg.ID, arg.IDTenant)
	var i SubgrupoProduto
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.IDGrupo,
		&i.Ativo,
		&i.IDTenant,
	)
	return i, err
}

const getSubGrupoProdutoByNome = `-- name: GetSubGrupoProdutoByNome :one
select id, nome, id_grupo, ativo, id_tenant from public.subgrupo_produtos where nome = $1 and id_tenant = $2 and ativo = true
`

type GetSubGrupoProdutoByNomeParams struct {
	Nome     string    `json:"nome"`
	IDTenant uuid.UUID `json:"id_tenant"`
}

func (q *Queries) GetSubGrupoProdutoByNome(ctx context.Context, arg GetSubGrupoProdutoByNomeParams) (SubgrupoProduto, error) {
	row := q.db.QueryRow(ctx, getSubGrupoProdutoByNome, arg.Nome, arg.IDTenant)
	var i SubgrupoProduto
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.IDGrupo,
		&i.Ativo,
		&i.IDTenant,
	)
	return i, err
}

const insertSubGrupoProduto = `-- name: InsertSubGrupoProduto :one
insert into public.subgrupo_produtos (nome, id_grupo, ativo, id_tenant) values ($1, $2, $3, $4) returning id, nome, id_grupo, ativo, id_tenant
`

type InsertSubGrupoProdutoParams struct {
	Nome     string    `json:"nome"`
	IDGrupo  uuid.UUID `json:"id_grupo"`
	Ativo    bool      `json:"ativo"`
	IDTenant uuid.UUID `json:"id_tenant"`
}

func (q *Queries) InsertSubGrupoProduto(ctx context.Context, arg InsertSubGrupoProdutoParams) (SubgrupoProduto, error) {
	row := q.db.QueryRow(ctx, insertSubGrupoProduto,
		arg.Nome,
		arg.IDGrupo,
		arg.Ativo,
		arg.IDTenant,
	)
	var i SubgrupoProduto
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.IDGrupo,
		&i.Ativo,
		&i.IDTenant,
	)
	return i, err
}

const updateSubGrupoProduto = `-- name: UpdateSubGrupoProduto :one
update public.subgrupo_produtos set nome = $1, ativo = $2 where id = $3 and id_tenant = $4 returning id, nome, id_grupo, ativo, id_tenant
`

type UpdateSubGrupoProdutoParams struct {
	Nome     string    `json:"nome"`
	Ativo    bool      `json:"ativo"`
	ID       uuid.UUID `json:"id"`
	IDTenant uuid.UUID `json:"id_tenant"`
}

func (q *Queries) UpdateSubGrupoProduto(ctx context.Context, arg UpdateSubGrupoProdutoParams) (SubgrupoProduto, error) {
	row := q.db.QueryRow(ctx, updateSubGrupoProduto,
		arg.Nome,
		arg.Ativo,
		arg.ID,
		arg.IDTenant,
	)
	var i SubgrupoProduto
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.IDGrupo,
		&i.Ativo,
		&i.IDTenant,
	)
	return i, err
}
