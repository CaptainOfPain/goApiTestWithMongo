package extensions

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/golobby/container"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"test.com/apiTest/configurations"
	"test.com/apiTest/repositories"
	"test.com/apiTest/services"
)

func RegisterConfiguration() {
	container.Singleton(func() configurations.Configuration {
		jsonFile, err := os.Open("config.json")
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var config = &configurations.Configuration{}
		json.Unmarshal(byteValue, config)

		if err != nil {
			log.Fatal(err)
		}
		return *config
	})
}

func RegisterMongoDriver() {
	container.Transient(func() mongo.Database {
		var config = configurations.Configuration{}
		container.Make(&config)
		client, _ := mongo.NewClient(options.Client().ApplyURI(config.ConnectionString))
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		err := client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		return *client.Database(config.Database)
	})
}

func RegisterRepositories() {
	container.Transient(func() repositories.UsersRepository {
		repo := repositories.UsersMongoRepository{}
		container.Make(&repo.Database)
		return repo
	})
	container.Transient(func() repositories.TasksRepository {
		repo := repositories.TasksMongoRepository{}
		container.Make(&repo.Database)
		return repo
	})
}

func RegisterServices() {
	container.Transient(func() services.UsersService {
		var repository repositories.UsersRepository
		container.Make(&repository)

		service := &services.UsersServiceImplementation{}
		service.AddRepository(repository)

		return service
	})

	container.Transient(func() services.SignInService {
		service := &services.SignInSeviceImplementation{}
		container.Make(&service.Config)
		container.Make(&service.UsersRepository)

		return service
	})

	container.Transient(func() services.TasksService {
		service := &services.TasksServiceImplementation{}
		container.Make(&service.Repository)

		return service
	})
}
