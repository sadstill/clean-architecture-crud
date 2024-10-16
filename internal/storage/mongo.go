package storage

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Username     string        `bson:"username"`
	PasswordHash string        `bson:"password"`
	Email        string        `bson:"email"`
}
