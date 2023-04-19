// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUsuario = `-- name: CreateUsuario :one
insert into usuarios (email, senha, id_tenant)
values ($1, $2, $3) returning id, email, senha, id_tenant
`

type CreateUsuarioParams struct {
	Email    string    `json:"email"`
	Senha    string    `json:"senha"`
	IDTenant uuid.UUID `json:"id_tenant"`
}

func (q *Queries) CreateUsuario(ctx context.Context, arg CreateUsuarioParams) (Usuario, error) {
	row := q.db.QueryRow(ctx, createUsuario, arg.Email, arg.Senha, arg.IDTenant)
	var i Usuario
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Senha,
		&i.IDTenant,
	)
	return i, err
}

const selectUsuarioPorEmail = `-- name: SelectUsuarioPorEmail :one
select id, email, senha, id_tenant from usuarios where email = ($1)
`

func (q *Queries) SelectUsuarioPorEmail(ctx context.Context, email string) (Usuario, error) {
	row := q.db.QueryRow(ctx, selectUsuarioPorEmail, email)
	var i Usuario
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Senha,
		&i.IDTenant,
	)
	return i, err
}
