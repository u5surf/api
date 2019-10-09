package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ScreenShot struct {
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Filename string `json:"filename"`
}

type Ticket struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	URL        string     `json:"url"`
	Screenshot ScreenShot `json:"screenshot"`
}

func main() {
	handler := mux.NewRouter()                                       // create router for handling api endpoints
	handler.HandleFunc("/api/honeyclient/create", CreateHoneyClient) // POST endpoint for creating honeyclient
	handler.HandleFunc("/api/honeyclient/{id}", GetHoneyClientById)  // GET endpoint for getting honeyclient by id
	fmt.Println("Server starting on localhost:8080...")
	http.ListenAndServe(":8080", handler) // listen to requests on localhost:8080
}

func CreateHoneyClient(w http.ResponseWriter, r *http.Request) {
	var ticket Ticket
	json.NewDecoder(r.Body).Decode(&ticket)    // decodes json from request into Ticket struct
	ticket.ID, ticket.Name = 1, "Dummy ticket" // ID and Name are not part of the request
	fmt.Fprintf(w, "Ticket '%s' (%d) will investigate url '%s' with screenshot dimensions %d x %d", ticket.Name, ticket.ID, ticket.URL, ticket.Screenshot.Width, ticket.Screenshot.Height)
}

func GetHoneyClientById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)               // get dynamic variables from mux handler
	id, _ := strconv.Atoi(vars["id"]) // get integer "ID" from vars
	ticket := Ticket{ID: id, Name: "Hardcoded ticket", URL: "https://example.com"}
	fmt.Fprintf(w, "Fetched ticket '%s' from ID '%d' with URL '%s'", ticket.Name, ticket.ID, ticket.URL)
}
