// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type GrupoProduto struct {
	ID       uuid.UUID   `json:"id"`
	Nome     string      `json:"nome"`
	IDTenant uuid.UUID   `json:"id_tenant"`
	Ativo    pgtype.Bool `json:"ativo"`
}

type Tenant struct {
	ID uuid.UUID `json:"id"`
}

type Usuario struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Senha    string    `json:"senha"`
	IDTenant uuid.UUID `json:"id_tenant"`
}
