package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/golobby/container"
	"github.com/gorilla/mux"
	"test.com/apiTest/helpers"
	"test.com/apiTest/services"
	viewmodels "test.com/apiTest/viewModels"
)

//ApplyUserRoutes for router
func ApplyUserRoutes(router *mux.Router) {
	router.HandleFunc("/api/users/", getUsers).Methods("GET")
	router.HandleFunc("/api/users/", addUser).Methods("POST")
}

func getUsers(writer http.ResponseWriter, request *http.Request) {
	var service services.UsersService
	container.Make(&service)

	users := service.GetUsers()

	helpers.JsonOK(writer, users)
}

func addUser(writer http.ResponseWriter, request *http.Request) {
	var service services.UsersService
	container.Make(&service)

	var viewModel viewmodels.AddUserViewModel
	json.NewDecoder(request.Body).Decode(&viewModel)

	service.AddUser(viewModel.UserName, viewModel.Email)

	helpers.JsonOK(writer, nil)
}
