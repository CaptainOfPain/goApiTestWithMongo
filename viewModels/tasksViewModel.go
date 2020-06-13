package viewmodels

import "time"

type TaskViewModel struct {
	Id             string
	Name           string
	OwnerUserId    string
	AssignedUserId *string
	CreatedDate    time.Time
	IsActive       bool
}

type TaskDetailedViewModel struct {
	Id             string
	Name           string
	OwnerUserId    string
	AssignedUserId *string
	CreatedDate    time.Time
	MinutesSpent   int
	Content        string
	IsActive       bool
}

type AddTaskViewModel struct {
	Name    string
	Content string
}

type UpdateTaskViewModel struct {
	TaskId  string
	Name    string
	Content string
}

type AssignUserToTaskViewModel struct {
	TaskId string
	UserId string
}

type ArchiveTaskViewModel struct {
	TaskId string
}
