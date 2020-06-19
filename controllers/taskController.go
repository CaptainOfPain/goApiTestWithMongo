package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/beevik/guid"
	"github.com/golobby/container"
	"github.com/gorilla/mux"
	"test.com/apiTest/helpers"
	"test.com/apiTest/services"
	viewmodels "test.com/apiTest/viewModels"
)

func ApplyTasksRoutes(router *mux.Router) {
	router.Handle("/api/tasks/", helpers.AuthenticationMiddleware(http.Handler(helpers.RootHandler(addTask)))).Methods("POST")
	router.Handle("/api/tasks/", helpers.AuthenticationMiddleware(http.Handler(helpers.RootHandler(getTasks)))).Methods("GET")
	router.Handle("/api/tasks/", helpers.AuthenticationMiddleware(http.Handler(helpers.RootHandler(archiveTask)))).Methods("DELETE")
	router.Handle("/api/tasks/", helpers.AuthenticationMiddleware(http.Handler(helpers.RootHandler(updateTask)))).Methods("PUT")
	router.Handle("/api/tasks/{taskId}", helpers.AuthenticationMiddleware(http.Handler(helpers.RootHandler(getTask)))).Methods("GET")
}

func addTask(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
	var service services.TasksService
	container.Make(&service)

	var viewModel viewmodels.AddTaskViewModel
	json.NewDecoder(request.Body).Decode(&viewModel)
	userId, err := guid.ParseString(request.Header.Get("userId"))
	err = service.AddTask(*guid.New(), viewModel.Name, *userId, viewModel.Content)

	return nil, err
}

func getTasks(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
	var service services.TasksService
	container.Make(&service)

	userId, parserError := guid.ParseString(request.Header.Get("userId"))
	if parserError != nil {
		return nil, parserError
	}
	users, err := service.GetTasks(*userId)
	if err != nil {
		return nil, err
	}
	tasksViewModels := []viewmodels.TaskViewModel{}
	for _, value := range users {
		tasksViewModels = append(tasksViewModels, viewmodels.TaskViewModel{
			Id:             value.Id,
			AssignedUserId: value.AssignedUserId,
			Name:           value.Name,
			CreatedDate:    value.CreatedDate,
			IsActive:       value.IsActive,
			OwnerUserId:    value.OwnerUserId,
		})
	}
	return users, err
}

func getTask(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
	var service services.TasksService
	container.Make(&service)

	vars := mux.Vars(request)
	taskId, _ := guid.ParseString(vars["taskId"])

	task, err := service.GetTask(*taskId)
	if err != nil {
		return nil, err
	}

	vm := viewmodels.TaskDetailedViewModel{
		Id:             task.Id,
		AssignedUserId: task.AssignedUserId,
		Name:           task.Name,
		CreatedDate:    task.CreatedDate,
		IsActive:       task.IsActive,
		OwnerUserId:    task.OwnerUserId,
		Content:        task.Content,
		MinutesSpent:   task.MinutesSpent,
	}
	return vm, err
}

func archiveTask(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
	var service services.TasksService
	container.Make(&service)

	userId, parserError := guid.ParseString(request.Header.Get("userId"))
	if parserError != nil {
		return nil, parserError
	}

	var viewModel viewmodels.ArchiveTaskViewModel
	json.NewDecoder(request.Body).Decode(&viewModel)
	taskId, _ := guid.ParseString(viewModel.TaskId)
	err := service.ArchiveTask(*taskId, *userId)

	return nil, err
}

func updateTask(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
	var service services.TasksService
	container.Make(&service)

	userId, parserError := guid.ParseString(request.Header.Get("userId"))
	if parserError != nil {
		return nil, parserError
	}

	var viewModel viewmodels.UpdateTaskViewModel
	json.NewDecoder(request.Body).Decode(&viewModel)
	taskId, _ := guid.ParseString(viewModel.TaskId)

	err := service.UpdateTask(*taskId, viewModel.Name, viewModel.Content, *userId)

	return nil, err
}

func assignUserToTask(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
	var service services.TasksService
	container.Make(&service)

	currentUserId, parserError := guid.ParseString(request.Header.Get("userId"))
	if parserError != nil {
		return nil, parserError
	}

	var viewModel viewmodels.AssignUserToTaskViewModel
	json.NewDecoder(request.Body).Decode(&viewModel)
	taskId, _ := guid.ParseString(viewModel.TaskId)
	userId, _ := guid.ParseString(viewModel.UserId)

	err := service.AssignUserToTask(*taskId, *userId, *currentUserId)

	return nil, err
}
