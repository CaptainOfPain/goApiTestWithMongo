package services

import (
	"github.com/beevik/guid"
	"test.com/apiTest/models"
	"test.com/apiTest/repositories"
)

type UsersService interface {
	AddUser(userName string, email string, password string) error
	GetUsers() ([]models.User, error)
	GetUser(id guid.Guid) (models.User, error)
}

type UsersServiceImplementation struct {
	repository repositories.UsersRepository
}

//AddUser creates user in repository
func (service UsersServiceImplementation) AddUser(userName string, email string, password string) error {
	user := models.CreateUser(guid.New().String(), userName, email, password)
	return service.repository.Add(*user)
}

//GetUsers get users from repository
func (service UsersServiceImplementation) GetUsers() ([]models.User, error) {
	return service.repository.Browse()
}

func (service UsersServiceImplementation) GetUser(id guid.Guid) (models.User, error) {
	return service.repository.Get(id.String())
}

func (service *UsersServiceImplementation) AddRepository(repository repositories.UsersRepository) {
	service.repository = repository
}
