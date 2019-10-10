package routes

import (
	"encoding/json"
	"fmt"
	"github.com/csci4950tgt/api/models"
	"net/http"
)

func CreateHoneyClient(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	json.NewDecoder(r.Body).Decode(&ticket)    // decodes json from request into Ticket struct
	ticket.ID, ticket.Name = 1, "Dummy ticket" // ID and Name are not part of the request
	fmt.Fprintf(w, "Ticket '%s' (%d) will investigate url '%s' with screenshot dimensions %d x %d", ticket.Name, ticket.ID, ticket.URL, ticket.Screenshot.Width, ticket.Screenshot.Height)
}
