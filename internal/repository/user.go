package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"rest-api-crud/internal/apperror"
	"rest-api-crud/internal/converter"
	"rest-api-crud/internal/model"
	"rest-api-crud/internal/storage"
	"rest-api-crud/pkg/logging"
)

var _ UserRepository = (*userRepo)(nil)

type userRepo struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func NewUserRepo(database *mongo.Database, collection string, logger *logging.Logger) UserRepository {
	return &userRepo{
		collection: database.Collection(collection),
		logger:     logger,
	}
}

func (r *userRepo) Create(ctx context.Context, user model.User) (string, error) {
	userMongo, err := converter.ToUserMongo(user)
	if err != nil {
		return "", err
	}

	result, err := r.collection.InsertOne(ctx, userMongo)
	if err != nil {
		return "", fmt.Errorf("failed to create users due to apperror %v", err)
	}

	r.logger.Debug("Converting InsertId to ObjectId")
	oid, ok := result.InsertedID.(bson.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	r.logger.Trace(user)
	return "", fmt.Errorf("failed to convert objectID to hex")
}

func (r *userRepo) FindById(ctx context.Context, id string) (u model.User, err error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to objectID")
	}

	filter := bson.M{"_id": oid}

	result := r.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return u, apperror.NotFound
		}
		return u, fmt.Errorf("failed to find users by id: %s due to err %s", id, result.Err())
	}

	var userMongo *storage.UserMongo
	if err = result.Decode(&userMongo); err != nil {
		return u, fmt.Errorf("failed to decode users from DB due to apperror: %v", err)
	}

	u = converter.ToUser(*userMongo)

	return u, nil
}

func (r *userRepo) FindAll(ctx context.Context) (u []model.User, err error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return u, fmt.Errorf("failed to find all users due to apperror: %v", err)
	}

	var mongoUsers *[]storage.UserMongo
	if err = cursor.All(ctx, &mongoUsers); err != nil {
		return u, fmt.Errorf("failed to read all documents from cursor. apperror : %v", err)
	}

	u = converter.ToUserSlice(*mongoUsers)

	return u, nil
}

func (r *userRepo) Update(ctx context.Context, user model.User) error {
	userMongo, err := converter.ToUserMongo(user)
	if err != nil {
		return err
	}

	userBytes, err := bson.Marshal(userMongo)
	if err != nil {
		return fmt.Errorf("failed to marshal users, apperror : %v", err)
	}

	var updatedUserObj bson.M
	err = bson.Unmarshal(userBytes, &updatedUserObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshall users bytes, apperror : %v", err)
	}

	delete(updatedUserObj, "_id")

	filter := bson.M{"_id": userMongo.ID}
	update := bson.M{"$set": updatedUserObj}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute update users query. apperror: %v", err)
	}

	if result.MatchedCount == 0 {
		return apperror.NotFound
	}

	r.logger.Tracef("Matched %d documents and Modified %d documents", result.MatchedCount, result.ModifiedCount)

	return nil
}

func (r *userRepo) DeleteById(ctx context.Context, id string) error {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert users ID to ObjectID. ID = %s", id)
	}

	filter := bson.M{"_id": oid}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. apperror: %v", err)
	}

	if result.DeletedCount == 0 {
		return apperror.NotFound
	}

	r.logger.Tracef("Deleted %d documents", result.DeletedCount)

	return nil
}
