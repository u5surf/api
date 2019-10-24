package routes

import (
	"encoding/json"
	"fmt"
	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
	"io/ioutil"
	"net/http"
	"os/exec"
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
	ticket.ID = 1

	// Create and write to json file
	file_name := fmt.Sprintf("../honeyclient/input/%d.json", ticket.ID)
	file, _ := json.MarshalIndent(ticket, "", " ")
	err := ioutil.WriteFile(file_name, file, 0755)

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusInternalServerError, "Failed to write file for honeyclient to consume.")

		fmt.Println("Failed to write file for honeyclient to consume:")
		fmt.Println(err)

		return
	}

	// Run honeyclient and return error if failed
	exeCmd := exec.Command("npm", "start")
	exeCmd.Dir = "../honeyclient/"
	err = exeCmd.Run()
	// TODO: we may need to parse the byte[] output to string
	// output, err = exeCmd.Output()

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusInternalServerError, "Failed to start the honeyclient.")

		fmt.Println("Failed to start honeyclient: ")
		fmt.Println(err)

		return
	} else {
		info := fmt.Sprintf("The output at honeyclient/output/%d", ticket.ID)
		fmt.Println(info)

	}

	// Initialize Response
	msg := fmt.Sprintf("Create ticket at honeyclient/output/%d named '%s' with ID '%d' and URL '%s'", ticket.ID, ticket.Name, ticket.ID, ticket.URL)
	res := models.Response{
		Success: true,
		Message: &msg,
	}

	util.WriteHttpResponse(w, res)
}
