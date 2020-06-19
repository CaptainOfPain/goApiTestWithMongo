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

//ApplyUserRoutes for router
func ApplyUserRoutes(router *mux.Router) {
	router.Handle("/api/users/", helpers.AuthenticationMiddleware(helpers.RootHandler(getUsers))).Methods("GET")
	router.Handle("/api/users/", helpers.RootHandler(addUser)).Methods("POST")
	router.Handle("/api/users/{userId}", helpers.AuthenticationMiddleware(helpers.RootHandler(getUser))).Methods("GET")
	router.Handle("/api/users/signIn/", helpers.RootHandler(signIn)).Methods("POST")
}

func getUsers(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
	var service services.UsersService
	container.Make(&service)

	users, err := service.GetUsers()

	return users, err
}

func addUser(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
	var service services.UsersService
	container.Make(&service)

	var viewModel viewmodels.AddUserViewModel
	json.NewDecoder(request.Body).Decode(&viewModel)

	err := service.AddUser(viewModel.UserName, viewModel.Email, viewModel.Password)

	return nil, err
}

func getUser(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
	var service services.UsersService
	container.Make(&service)

	vars := mux.Vars(request)
	userId, _ := guid.ParseString(vars["userId"])

	return service.GetUser(*userId)
}

func signIn(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
	var service services.SignInService
	container.Make(&service)

	var viewModel viewmodels.SignInViewModel
	json.NewDecoder(request.Body).Decode(&viewModel)

	return service.SignIn(viewModel.UserName, viewModel.Password)
}
