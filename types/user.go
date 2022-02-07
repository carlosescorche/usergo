package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName string             `bson:"firstname,omitempty" json:"firstname"`
	LastName  string             `bson:"lastname,omitempty" json:"lastname"`
	Username  string             `bson:"username,omitempty" json:"username"`
	Email     string             `bson:"email,omitempty" json:"email"`
	Password  string             `bson:"password,omitempty" json:"password"`
}
