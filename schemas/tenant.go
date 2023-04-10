package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tenant struct {
	Id_tenant primitive.ObjectID `json:"_IdTenant" bson:"_IdTenant"`
}
