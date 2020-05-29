package services

import (
	"github.com/beevik/guid"
	"github.com/golobby/container"
	"test.com/apiTest/models"
	"test.com/apiTest/repositories"
)

type UsersService interface {
	AddUser(userName string, email string)
	GetUsers() []models.User
}

type UsersServiceImplementation struct {
	repository repositories.UsersRepository
}

func (service UsersServiceImplementation) AddUser(userName string, email string) {
	var repo repositories.UsersRepository
	container.Make(&repo)
	service.repository = repo

	user := models.CreateUser(guid.New().String(), userName, email)
	service.repository.Add(*user)
}

func (service UsersServiceImplementation) GetUsers() []models.User {
	var repo repositories.UsersRepository
	container.Make(&repo)
	service.repository = repo

	return service.repository.Browse()
}
