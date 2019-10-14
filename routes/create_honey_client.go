package routes

import (
	"encoding/json"
	"fmt"
	"github.com/csci4950tgt/api/models"
	"net/http"
	"io/ioutil"
)

func CreateHoneyClient(w http.ResponseWriter, r *http.Request) {
	// Create a new ticket for handling
	var ticket models.Ticket

	// decodes json from request into Ticket struct
	json.NewDecoder(r.Body).Decode(&ticket)

	// ID and Name are not part of the request
	ticket.ID, ticket.Name = 1, "Dummy ticket"

	// Create and write to json file
	file_name := fmt.Sprintf("../honeyclient/input/%d.json", ticket.ID)
	file, _ := json.MarshalIndent(ticket, "", " ")
	err := ioutil.WriteFile(file_name, file, 0755)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error happens in connecting honeyclient. %s.\n",
					err.Error())
	} else {
		fmt.Fprintf(w, "Ticket '%s' (%d) will investigate url '%s' with screenshot dimensions %d x %d", ticket.Name, ticket.ID, ticket.URL, ticket.Screenshot.Width, ticket.Screenshot.Height)
	}
}
