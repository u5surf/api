package main

import (
	"fmt"
	"net/http"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/routes"
	"github.com/gorilla/mux"
)

func main() {
	handler := mux.NewRouter()                                              // create router for handling api endpoints
	handler.HandleFunc("/api/honeyclient/create", routes.CreateHoneyClient) // POST endpoint for creating honeyclient
	handler.HandleFunc("/api/honeyclient/{id}", routes.GetHoneyClientById)  // GET endpoint for getting honeyclient by id

	models.InitDB("host=127.0.0.1 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")

	fmt.Println("Server starting...")
	http.ListenAndServe(":8080", handler) // listen to requests on localhost:8080
}
