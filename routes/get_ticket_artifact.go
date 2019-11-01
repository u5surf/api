package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
	"github.com/gorilla/mux"
)

func GetTicketArtifact(w http.ResponseWriter, r *http.Request) {
	// Get variables from route handler
	vars := mux.Vars(r) // get dynamic variables from mux handler
	artifact := vars["artifact"]
	ticketId, err := strconv.Atoi(vars["id"]) // get integer "id" from vars

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusBadRequest, "Missing required parameter: id.")

		return
	}

	// TODO: Actually get desired artifact from ticket

	// Initialize Response
	msg := fmt.Sprintf("Not yet implemented... ticketId: %d, Artifact: %s", ticketId, artifact)
	res := models.Response{
		Success: true,
		Message: &msg,
		// TODO: Send artifact in response
	}

	util.WriteHttpResponse(w, res)
}
