package repositorios

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"romaControle/schemas"
)

type UsuarioRepositorio struct {
	db *mongo.Database
}

func NovoRepUser(db *mongo.Database) *UsuarioRepositorio {
	return &UsuarioRepositorio{db: db}
}

func (repositorio *UsuarioRepositorio) CadastrarUsuario(ctx context.Context, usuario schemas.Usuario) (*schemas.Usuario, error) {

	var usuario_inserido *schemas.Usuario
	usuario.Tenant = schemas.Tenant{Id_tenant: primitive.NewObjectID()}

	err := validator.New().Struct(usuario)
	if err != nil {
		return nil, err
	}

	id, e := repositorio.db.Collection("usuario").InsertOne(ctx, usuario)
	if e != nil {
		return nil, fmt.Errorf("Erro ao inserir o cadastro de usuario")
	}

	erro := repositorio.db.Collection("usuario").FindOne(ctx, bson.M{"_id": id.InsertedID}).Decode(&usuario_inserido)
	if erro != nil {
		return nil, erro
	}

	return usuario_inserido, nil

}
