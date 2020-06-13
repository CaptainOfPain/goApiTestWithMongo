package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"test.com/apiTest/models"
)

type TasksRepository interface {
	Add(task models.Task) error
	Update(task models.Task) error
	Remove(task models.Task) error
	Get(id string) (models.Task, error)
	Browse() ([]models.Task, error)
}

type TasksMongoRepository struct {
	Database mongo.Database
}

func (repo TasksMongoRepository) Add(task models.Task) error {
	collection := repo.Database.Collection("Tasks")

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	_, err := collection.InsertOne(ctx, task)
	return err
}

func (repo TasksMongoRepository) Update(task models.Task) error {
	collection := repo.Database.Collection("Tasks")
	filter := bson.D{{"id", task.Id}}
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	_, err := collection.UpdateOne(ctx, filter, task)
	return err
}

func (repo TasksMongoRepository) Remove(task models.Task) error {
	collection := repo.Database.Collection("Tasks")
	filter := bson.D{{"id", task.Id}}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := collection.DeleteOne(ctx, filter)
	return err
}

func (repo TasksMongoRepository) Get(id string) (models.Task, error) {
	collection := repo.Database.Collection("Tasks")
	filter := bson.D{{"id", id}}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var result models.Task
	err := collection.FindOne(ctx, filter).Decode(&result)
	return result, err
}

func (repo TasksMongoRepository) Browse() ([]models.Task, error) {
	collection := repo.Database.Collection("Tasks")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var result []models.Task
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	if result == nil {
		result = []models.Task{}
	}

	return result, err
}
