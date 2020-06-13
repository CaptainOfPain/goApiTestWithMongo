package services

import (
	"github.com/beevik/guid"
	"test.com/apiTest/models"
	"test.com/apiTest/repositories"
)

type UsersService interface {
	AddUser(userName string, email string, password string)
	GetUsers(c chan []models.User)
	GetUser(id guid.Guid, c chan models.User)
}

type UsersServiceImplementation struct {
	repository repositories.UsersRepository
}

//AddUser creates user in repository
func (service UsersServiceImplementation) AddUser(userName string, email string, password string) {
	user := models.CreateUser(guid.New().String(), userName, email, password)
	service.repository.Add(*user)
}

//GetUsers get users from repository
func (service UsersServiceImplementation) GetUsers(c chan []models.User) {
	service.repository.Browse(c)
}

func (service UsersServiceImplementation) GetUser(id guid.Guid, c chan models.User) {
	service.repository.Get(id.String(), c)
}

func (service *UsersServiceImplementation) AddRepository(repository repositories.UsersRepository) {
	service.repository = repository
}
