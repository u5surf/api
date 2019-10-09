package main

import (
	"fmt"
	"github.com/csci4950tgt/api/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	handler := mux.NewRouter()                                              // create router for handling api endpoints
	handler.HandleFunc("/api/honeyclient/create", routes.CreateHoneyClient) // POST endpoint for creating honeyclient
	handler.HandleFunc("/api/honeyclient/{id}", routes.GetHoneyClientById)  // GET endpoint for getting honeyclient by id
	fmt.Println("Server starting on localhost:8080...")
	http.ListenAndServe(":8080", handler) // listen to requests on localhost:8080
}
