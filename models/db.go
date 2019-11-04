package models

import (
	"fmt"
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = gorm.Open("postgres", dataSourceName)

	if err != nil {
		fmt.Println(err)

		panic("Failed to connect to database!")
	}

	// Migrate the schema
	fmt.Println("Setting up the database...")
	db.AutoMigrate(&Ticket{}, &ScreenShot{}, &FileArtifact{})
}

func Rows() (*sql.Rows, error) {
	// tickets_table := db.Find(&Ticket{})
	// fmt.Println(tickets_table)
	// row := tickets_table.Row()
	// fmt.Println(row)

	return db.Find(&Ticket{}).Rows()

}
