package services

import (
	"github.com/beevik/guid"
	"test.com/apiTest/models"
	"test.com/apiTest/repositories"
)

type TasksService interface {
	AddTask(id guid.Guid, name string, ownerUserId guid.Guid, content string) error
	UpdateTask(id guid.Guid, name string, content string, currentUserId guid.Guid) error
	AssignUserToTask(taskId guid.Guid, userId guid.Guid, currentUserId guid.Guid) error
	ArchiveTask(taskId guid.Guid, currentUserId guid.Guid) error
	GetTasks(currentUserId guid.Guid) ([]models.Task, error)
	GetTask(taskId guid.Guid) (models.Task, error)
}

type TasksServiceImplementation struct {
	Repository repositories.TasksRepository
}

func (service TasksServiceImplementation) AddTask(id guid.Guid, name string, ownerUserId guid.Guid, content string) error {
	task := models.CreateTask(id, name, ownerUserId, content)

	return service.Repository.Add(*task)
}

func (service TasksServiceImplementation) UpdateTask(id guid.Guid, name string, content string, currentUserId guid.Guid) error {
	task, err := service.Repository.Get(id.String())
	if err != nil {
		return err
	}

	err = task.UpdateTask(name, content, currentUserId)
	if err != nil {
		return err
	}
	return service.Repository.Update(task)
}

func (service TasksServiceImplementation) AssignUserToTask(taskId guid.Guid, userId guid.Guid, currentUserId guid.Guid) error {
	task, err := service.Repository.Get(taskId.String())
	if err != nil {
		return err
	}

	err = task.AssignUser(currentUserId)
	if err != nil {
		return err
	}
	return service.Repository.Update(task)
}

func (service TasksServiceImplementation) ArchiveTask(taskId guid.Guid, currentUserId guid.Guid) error {
	task, err := service.Repository.Get(taskId.String())
	if err != nil {
		return err
	}

	err = task.Archive(currentUserId)
	if err != nil {
		return err
	}
	return service.Repository.Update(task)
}

func (service TasksServiceImplementation) GetTasks(currentUserId guid.Guid) ([]models.Task, error) {
	return service.Repository.Browse()
}

func (service TasksServiceImplementation) GetTask(taskId guid.Guid) (models.Task, error) {
	return service.Repository.Get(taskId.String())
}
