package user

import "go.mongodb.org/mongo-driver/v2/bson"

type UserMongo struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Username     string        `bson:"username"`
	PasswordHash string        `bson:"password"`
	Email        string        `bson:"email"`
}

type UserPG struct {
	ID           int64  `db:"_id,omitempty"`
	Username     string `db:"username"`
	PasswordHash string `db:"password"`
	Email        string `db:"email"`
}
