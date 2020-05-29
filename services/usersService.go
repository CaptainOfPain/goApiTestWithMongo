package services

import (
	"github.com/beevik/guid"
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
	user := models.CreateUser(guid.New().String(), userName, email)
	service.repository.Add(*user)
}

func (service UsersServiceImplementation) GetUsers() []models.User {
	return service.repository.Browse()
}

func (service *UsersServiceImplementation) AddRepository(repository repositories.UsersRepository) {
	service.repository = repository
}
