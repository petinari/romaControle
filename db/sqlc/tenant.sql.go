// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: tenant.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createTenant = `-- name: CreateTenant :one
insert into tenant default values returning id
`

func (q *Queries) CreateTenant(ctx context.Context) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createTenant)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const selectTenantJoinUsuario = `-- name: SelectTenantJoinUsuario :many
SELECT T.id , u.email
	FROM public.tenant as T inner join public.usuarios as u on id_tenant = T.id
`

type SelectTenantJoinUsuarioRow struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

func (q *Queries) SelectTenantJoinUsuario(ctx context.Context) ([]SelectTenantJoinUsuarioRow, error) {
	rows, err := q.db.Query(ctx, selectTenantJoinUsuario)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectTenantJoinUsuarioRow{}
	for rows.Next() {
		var i SelectTenantJoinUsuarioRow
		if err := rows.Scan(&i.ID, &i.Email); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
