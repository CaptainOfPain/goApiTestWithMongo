package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"test.com/apiTest/controllers"
	"test.com/apiTest/extensions"
)

func handleRequests() {
	router := mux.NewRouter()

	controllers.ApplyUserRoutes(router)

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	extensions.RegisterConfiguration()
	extensions.RegisterMongoDriver()
	extensions.RegisterRepositories()
	extensions.RegisterServices()

	handleRequests()
}
