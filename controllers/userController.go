package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/beevik/guid"
	"github.com/golobby/container"
	"github.com/gorilla/mux"
	"test.com/apiTest/helpers"
	"test.com/apiTest/models"
	"test.com/apiTest/services"
	viewmodels "test.com/apiTest/viewModels"
)

//ApplyUserRoutes for router
func ApplyUserRoutes(router *mux.Router) {
	test := http.HandlerFunc(getUsers)
	router.Handle("/api/users/", helpers.AuthenticationMiddleware(test)).Methods("GET")
	router.HandleFunc("/api/users/", addUser).Methods("POST")
	router.HandleFunc("/api/users/{userId}", getUser).Methods("GET")
	router.HandleFunc("/api/users/signIn", signIn).Methods("POST")
}

func getUsers(writer http.ResponseWriter, request *http.Request) {
	var service services.UsersService
	container.Make(&service)

	c := make(chan []models.User)
	go service.GetUsers(c)
	users := <-c

	helpers.JsonOK(writer, users)
}

func addUser(writer http.ResponseWriter, request *http.Request) {
	var service services.UsersService
	container.Make(&service)

	var viewModel viewmodels.AddUserViewModel
	json.NewDecoder(request.Body).Decode(&viewModel)

	service.AddUser(viewModel.UserName, viewModel.Email, viewModel.Password)

	helpers.JsonOK(writer, nil)
}

func getUser(writer http.ResponseWriter, request *http.Request) {
	var service services.UsersService
	container.Make(&service)

	c := make(chan models.User)

	vars := mux.Vars(request)
	userId, _ := guid.ParseString(vars["userId"])

	go service.GetUser(*userId, c)

	result := <-c

	helpers.JsonOK(writer, result)
}

func signIn(writer http.ResponseWriter, request *http.Request) {
	var service services.SignInService
	container.Make(&service)

	var viewModel viewmodels.SignInViewModel
	json.NewDecoder(request.Body).Decode(&viewModel)

	c := make(chan services.UserSignedInDto)
	go service.SignIn(viewModel.UserName, viewModel.Password, c)

	user := <-c
	helpers.JsonOK(writer, viewmodels.UserSignedInViewModel{Id: user.Id, Email: user.Email, UserName: user.UserName, JwtToken: user.JwtToken})
}
