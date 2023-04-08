package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Tenant Tenant             `json:"tenant" bson:"tenant"`
	Email  string             `json:"email" validate:"required,email"`
	Senha  string             `json:"senha" validate:"required"`
}
