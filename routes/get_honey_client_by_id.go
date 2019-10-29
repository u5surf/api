package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
	"github.com/gorilla/mux"
)

func GetHoneyClientById(w http.ResponseWriter, r *http.Request) {
	// If request type is not current request, return error.
	if (*r).Method != "GET" {
		msg := fmt.Sprintf("Method \"%s\" is not allowed.", (*r).Method)
		util.WriteHttpErrorCode(w, http.StatusMethodNotAllowed, msg)

		return
	}

	// Initialize header
	header := http.Header{}
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "GET")

	// Set header
	util.SetHeader(w, header)

	// Initialize Ticket struct
	vars := mux.Vars(r)                 // get dynamic variables from mux handler
	id, err := strconv.Atoi(vars["id"]) // get integer "ID" from var

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusBadRequest, "Missing required parameter: id.")

		return
	}

	ticket, err := models.GetTicketById(uint(id))

	if err != nil {
		msg := fmt.Sprintf("Failed to find ticket with ID %d.", id)
		util.WriteHttpErrorCode(w, http.StatusNotFound, msg)

		return
	}

	// Initialize Response
	res := models.Response{
		Success: true,
		Ticket:  ticket,
	}

	util.WriteHttpResponse(w, res)
}
