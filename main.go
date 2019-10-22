package main

import (
	"fmt"
	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/routes"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

func main() {
	handler := mux.NewRouter()                                              // create router for handling api endpoints
	handler.HandleFunc("/api/honeyclient/create", routes.CreateHoneyClient) // POST endpoint for creating honeyclient
	handler.HandleFunc("/api/honeyclient/{id}", routes.GetHoneyClientById)  // GET endpoint for getting honeyclient by id

	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")
	if err != nil {
		fmt.Println(err)

		panic("Failed to connect to database!")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.Ticket{})

	// Create
	db.Create(&models.Ticket{
		Name:       "Ticket 1234",
		URL:        "https://tdr.moe",
		Screenshot: models.ScreenShot{Width: 1280, Height: 720, Filename: "screenshot.png"},
	})

	// Read
	var ticket models.Ticket
	db.First(&ticket, 1234) // find ticket with id 1234
	// db.First(&ticket, "url = ?", "https://tdr.moe") // find ticket with url https://tdr.moe

	// Update
	db.Model(&ticket).Update("URL", "https://tdr.moe/2")

	// Delete
	db.Delete(&ticket)

	fmt.Println("Server starting...")
	http.ListenAndServe(":8080", handler) // listen to requests on localhost:8080
}
