package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
	"github.com/gorilla/mux"
)

func MethodNotAllowedHelper(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Method \"%s\" is not allowed.", (*r).Method)
	util.WriteHttpErrorCode(w, http.StatusMethodNotAllowed, msg)
}

func GetTickets(w http.ResponseWriter, r *http.Request) {
	// If request type is not current request, return error.
	if (*r).Method != "GET" {
		MethodNotAllowedHelper(w, r)
		return
	}

	// Initialize header
	header := http.Header{}
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "GET")

	// Set header
	util.SetHeader(w, header)

	// TODO: Actually get all tickets

	// Initialize Response
	res := models.Response{
		Success: true,
		// TODO: Return tickets in response
	}

	util.WriteHttpResponse(w, res)
}
