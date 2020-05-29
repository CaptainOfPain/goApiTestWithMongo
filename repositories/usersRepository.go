package repositories

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"test.com/apiTest/models"
)

type UsersRepository interface {
	Add(user models.User)
	Update(user models.User)
	Remove(user models.User)
	Get(id string) models.User
	Browse() []models.User
}

type UsersMongoRepository struct {
	database mongo.Database
}

func (repo UsersMongoRepository) Add(user models.User) {
	collection := repo.database.Collection("Users")

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	result, err := collection.InsertOne(ctx, user)
	if err != nil && result.InsertedID == nil {
		log.Fatal(err)
	}
}

func (repo UsersMongoRepository) Update(user models.User) {
	collection := repo.database.Collection("Users")
	filter := bson.D{{"id", user.Id}}
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	result, err := collection.UpdateOne(ctx, filter, user)
	if err != nil && result == nil {
		log.Fatal(err)
	}
}

func (repo UsersMongoRepository) Remove(user models.User) {
	collection := repo.database.Collection("Users")
	filter := bson.D{{"id", user.Id}}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil && result == nil {
		log.Fatal(err)
	}
}

func (repo UsersMongoRepository) Get(id string) models.User {
	collection := repo.database.Collection("Users")
	filter := bson.D{{"id", id}}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var result models.User
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == nil {
		log.Fatal(err)
	}

	return result
}

func (repo UsersMongoRepository) Browse() []models.User {
	collection := repo.database.Collection("Users")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var result []models.User
	if err = cursor.All(ctx, &result); err != nil {
		log.Fatal(err)
	}

	if result == nil {
		result = []models.User{}
	}

	return result
}

func (repo *UsersMongoRepository) AddDatabase(database mongo.Database) {
	repo.database = database
}
