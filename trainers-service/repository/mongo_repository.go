package repository

import (
	"context"
	"errors"

	"github.com/patrickovm/go-microservices/trainers-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrTrainerNotFound = errors.New("trainer not found")

type MongoRepository struct {
	trainersCollection *mongo.Collection
}

func NewMongoRepository(collection *mongo.Collection) *MongoRepository {
	return &MongoRepository{
		trainersCollection: collection,
	}
}

func (m MongoRepository) Create(ctx context.Context, trainer *model.Trainer) (string, error) {
	result, err := m.trainersCollection.InsertOne(ctx, bson.M{
		"first_name":     trainer.FirstName,
		"last_name":      trainer.LastName,
		"age":            trainer.Age,
		"finished_games": trainer.FinishedGames,
	})
	if err != nil {
		return "", nil
	}

	trainerId := result.InsertedID.(primitive.ObjectID).String()

	return trainerId, nil
}

func (m MongoRepository) FindOne(ctx context.Context, id string) (*model.Trainer, error) {
	trainerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var data model.Trainer

	err = m.trainersCollection.FindOne(ctx, bson.M{"_id": trainerId}).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (m MongoRepository) List(ctx context.Context) (*[]model.Trainer, error) {
	var data []model.Trainer

	cursor, err := m.trainersCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var trainer model.Trainer
		cursor.Decode(&trainer)
		data = append(data, trainer)
	}
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (m MongoRepository) Update(ctx context.Context, trainer *model.Trainer) (*model.Trainer, error) {
	filter := bson.M{"_id": trainer.Id}

	update := bson.M{
		"$set": trainer,
	}

	result, err := m.trainersCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, ErrTrainerNotFound
	}

	if result.ModifiedCount == 0 {
		return nil, err
	}

	return trainer, nil
}
