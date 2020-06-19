package models

import (
	"errors"
	"time"

	"github.com/beevik/guid"
)

type Task struct {
	Id             string `bson: "_id"`
	Name           string
	OwnerUserId    string
	AssignedUserId *string
	CreatedDate    time.Time
	MinutesSpent   int
	Content        string
	IsActive       bool
}

func CreateTask(id guid.Guid, name string, ownerUserId guid.Guid, content string) *Task {
	task := &Task{
		Id:          id.String(),
		Name:        name,
		OwnerUserId: ownerUserId.String(),
		Content:     content,
		IsActive:    true,
		CreatedDate: time.Now(),
	}

	return task
}

func (task *Task) Archive(userId guid.Guid) error {
	if (task.AssignedUserId != nil && *task.AssignedUserId == userId.String()) || task.OwnerUserId == userId.String() {
		task.IsActive = false
		return nil
	} else {
		return errors.New("Not valid user to perform this action")
	}
}

func (task *Task) LogTime(minutes int) {
	task.MinutesSpent += minutes
}

func (task *Task) UpdateTask(name string, content string, userId guid.Guid) error {
	if *task.AssignedUserId == userId.String() || task.OwnerUserId == userId.String() {
		task.Name = name
		task.Content = content
	} else {
		return errors.New("Not valid user to perform this action")
	}
	return nil
}

func (task *Task) AssignUser(userId guid.Guid) error {
	if *task.AssignedUserId == userId.String() || task.OwnerUserId == userId.String() {
		var id = userId.String()
		task.AssignedUserId = &id
	} else {
		return errors.New("Not valid user to perform this action")
	}
	return nil
}
