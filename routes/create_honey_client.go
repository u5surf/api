package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
)

func CreateHoneyClient(w http.ResponseWriter, r *http.Request) {
	// If request type is not current request, return error.
	if (*r).Method != "POST" {
		msg := fmt.Sprintf("Method \"%s\" is not allowed.", (*r).Method)
		util.WriteHttpErrorCode(w, http.StatusMethodNotAllowed, msg)

		return
	}

	// Initialize header
	header := http.Header{}
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "POST")

	// Set header
	util.SetHeader(w, header)

	// Create a new ticket for handling, encode request into struct
	var ticket models.Ticket
	json.NewDecoder(r.Body).Decode(&ticket)

	// saves the ticket in the database:
	err := models.CreateTicket(ticket)

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusInternalServerError, "Failed to write file for honeyclient to consume.")

		fmt.Println("Failed to write file for honeyclient to consume:")
		fmt.Println(err)

		return
	}

	// Initialize Response
	msg := fmt.Sprintf("Create ticket '%s' with ID '%d' and URL '%s'", ticket.Name, ticket.ID, ticket.URL)
	res := models.Response{
		Success: true,
		Message: &msg,
	}

	util.WriteHttpResponse(w, res)
}
