package usuario_repositorio

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"romaControle/config"
	db "romaControle/db/sqlc"
	"romaControle/models/usuario"
)

func CadastrarUsuarioTenant(usuario db.Usuario) (*db.Usuario, error) {
	ctx := context.Background()
	queries := db.New(config.DB)
	tenant, err := queries.CreateTenant(ctx)
	usuarioInserido, err := queries.CreateUsuario(ctx, db.CreateUsuarioParams{
		Email:    usuario.Email,
		Senha:    usuario.Senha,
		IDTenant: uuid.Must(tenant, err),
	})
	if err != nil {
		return nil, err
	}
	return &usuarioInserido, nil
}

func SelectUsuarioPorEmail(login usuario.UsuarioLogin) (*db.Usuario, error) {
	ctx := context.Background()
	queries := db.New(config.DB)
	usuarioDb, err := queries.SelectUsuarioPorEmail(ctx, login.Email)
	if err != nil {
		return nil, errors.New("Usuário não encontrado.")
	}
	return &usuarioDb, nil

}
