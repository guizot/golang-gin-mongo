package model

import (
	"time"
)

//User
type User struct {
	ID        int       `bson:"id"`
	Name      string    `bson:"name"`
	Address   string    `bson:"address"`
	Age       int       `bson:"age"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

//Users
type Users []User
