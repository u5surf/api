package routes

import (
	// "fmt"
	"net/http"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
)

func GetTickets(w http.ResponseWriter, r *http.Request) {
	// TODO: Actually get all tickets

	// get all data from database
	// rows, err := db.QueryContext(ctx, "SELECT ALL FROM gorm")
	rows, err := models.Rows()
	defer rows.Close()

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusNoContent, err.Error())

		return
	}

	for rows.Next() {
		var ticket models.Ticket
		err = rows.Scan(&ticket)
		if err != nil {
			util.WriteHttpErrorCode(w, http.StatusPartialContent, err.Error())
		} else {
			res := models.Response{
				Success: true,
				Ticket:  &ticket,
			}

			util.WriteHttpResponse(w, res)
		}
	}
}
