package usuario

import (
	"github.com/google/uuid"
)

type UsuarioOut struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	TenantOut TenantOut `json:"tenant"`
}

type TenantOut struct {
	IDTenant uuid.UUID `json:"id_tenant"`
}
