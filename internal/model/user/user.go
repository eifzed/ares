package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID         primitive.ObjectID `bson:"_id,omitempty"`
	FirstName      string             `bson:"first_name"`
	LastName       string             `bson:"last_name"`
	Email          string             `bson:"email"`
	PhoneNumber    string             `bson:"phone_number"`
	Password       string             `bson:"-"`
	PasswordHashed string             `bson:"password_hashed"`
	BirthDate      time.Time          `bson:"birth_date"`
	Roles          []UserRole         `bson:"roles"`
}

type Token struct {
	JWT        string
	ValidUntil string
}

type UserRole struct {
	ID   int64  `bson:"id"`
	Name string `bson:"name"`
}

type LoginParams struct {
	Email    string
	Password string
}
