package repositories

import (
	"context"
	"log"
	"os"
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
		os.Exit(1)
	}
}

func (repo UsersMongoRepository) Update(user models.User) {

}

func (repo UsersMongoRepository) Remove(user models.User) {

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
