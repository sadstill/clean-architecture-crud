package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func New(ctx context.Context, host, port, username, password, database, authDB string) (db *mongo.Database, err error) {
	var mongodbURL string
	var isAuth bool
	if username == "" && password == "" {
		mongodbURL = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		isAuth = true
		mongodbURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	}

	clientOptions := options.
		Client().
		ApplyURI(mongodbURL)
	if isAuth {
		if authDB == "" {
			authDB = database
		}
		clientOptions.SetAuth(options.Credential{
			AuthSource: authDB,
			Username:   username,
			Password:   password,
		})
	}

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB due to apperror: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB due to apperror: %v", err)
	}

	return client.Database(database), nil
}
