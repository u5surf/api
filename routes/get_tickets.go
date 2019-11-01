package routes

import (
	"net/http"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
)

func GetTickets(w http.ResponseWriter, r *http.Request) {
	// TODO: Actually get all tickets

	// Initialize Response
	msg := "Not yet implemented!!"
	res := models.Response{
		Success: true,
		Message: &msg,
		// TODO: Return tickets in response
	}

	util.WriteHttpResponse(w, res)
}
