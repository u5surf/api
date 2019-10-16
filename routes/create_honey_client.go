package routes

import (
	"encoding/json"
	"fmt"
	"github.com/csci4950tgt/api/models"
	"net/http"
	"io/ioutil"
)

func SetupCreateResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Method", "POST")
}

func CreateHoneyClient(w http.ResponseWriter, r *http.Request) {

	// Handling requests.
	SetupCreateResponse(&w, r)
	// If request type is not current request, return error.
	if (*r).Method != "POST" {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method \"%s\" is not allowed.\n",(*r).Method)
		return
	}

	// Create a new ticket for handling
	var ticket models.Ticket

	// decodes json from request into Ticket struct
	json.NewDecoder(r.Body).Decode(&ticket)

	// ID and Name are not part of the request
	ticket.ID = 1

	// Create and write to json file
	file_name := fmt.Sprintf("../honeyclient/input/%d.json", ticket.ID)
	file, _ := json.MarshalIndent(ticket, "", " ")
	err := ioutil.WriteFile(file_name, file, 0755)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error happens in connecting honeyclient. %s.\n",
					err.Error())
	} else {
		fmt.Fprintf(w, "Create ticket '%s' with ID '%d' and URL '%s'", ticket.Name, ticket.ID, ticket.URL)
	}
}
