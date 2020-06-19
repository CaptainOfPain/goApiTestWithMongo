package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"test.com/apiTest/models"
)

type UsersRepository interface {
	Add(user models.User) error
	Update(user models.User) error
	Remove(user models.User) error
	Get(id string) (models.User, error)
	GetByUserName(username string) (models.User, error)
	Browse() ([]models.User, error)
}

type UsersMongoRepository struct {
	Database mongo.Database
}

func (repo UsersMongoRepository) Add(user models.User) error {
	collection := repo.Database.Collection("Users")

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (repo UsersMongoRepository) Update(user models.User) error {
	collection := repo.Database.Collection("Users")
	filter := bson.D{{"id", user.Id}}
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	_, err := collection.UpdateOne(ctx, filter, user)
	return err
}

func (repo UsersMongoRepository) Remove(user models.User) error {
	collection := repo.Database.Collection("Users")
	filter := bson.D{{"id", user.Id}}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := collection.DeleteOne(ctx, filter)
	return err
}

func (repo UsersMongoRepository) Get(id string) (models.User, error) {
	collection := repo.Database.Collection("Users")
	filter := bson.D{{"id", id}}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var result models.User
	err := collection.FindOne(ctx, filter).Decode(&result)

	return result, err
}

func (repo UsersMongoRepository) GetByUserName(username string) (models.User, error) {
	collection := repo.Database.Collection("Users")
	filter := bson.D{{"username", username}}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var result models.User
	err := collection.FindOne(ctx, filter).Decode(&result)

	return result, err
}

func (repo UsersMongoRepository) Browse() ([]models.User, error) {
	collection := repo.Database.Collection("Users")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var result []models.User
	err = cursor.All(ctx, &result)

	if result == nil {
		result = []models.User{}
	}

	return result, err
}

func (repo *UsersMongoRepository) AddDatabase(database mongo.Database) {
	repo.Database = database
}
