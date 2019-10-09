package routes

import (
	"fmt"
	"github.com/csci4950tgt/api/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetHoneyClientById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)               // get dynamic variables from mux handler
	id, _ := strconv.Atoi(vars["id"]) // get integer "ID" from vars
	ticket := models.Ticket{ID: id, Name: "Hardcoded ticket", URL: "https://example.com"}
	fmt.Fprintf(w, "Fetched ticket '%s' from ID '%d' with URL '%s'", ticket.Name, ticket.ID, ticket.URL)
}
