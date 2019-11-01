package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
	"github.com/gorilla/mux"
)

func GetTicketArtifacts(w http.ResponseWriter, r *http.Request) {
	// Get variables from route handler
	vars := mux.Vars(r)                 // get dynamic variables from mux handler
	id, err := strconv.Atoi(vars["id"]) // get integer "id" from vars

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusBadRequest, "Missing required parameter: id.")

		return
	}

	// TODO: Get artifacts for ticket from database

	// Initialize Response
	msg := fmt.Sprintf("Not yet implemented... Id: %d", id)
	res := models.Response{
		Success: true,
		Message: &msg,
		// TODO: add artifacts data to response
	}

	util.WriteHttpResponse(w, res)
}
