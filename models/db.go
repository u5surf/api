package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func InitDB(dataSourceName string) {
	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil {
		fmt.Println(err)

		panic("Failed to connect to database!")
	}
	defer db.Close()

	// Migrate the schema
	fmt.Println("Setting up the database...")
	db.AutoMigrate(&Ticket{})
}
