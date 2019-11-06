package models

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	ConnectToTesting()
}

func ConnectToTesting() {
	InitDB("host=127.0.0.1 port=5432 user=testing dbname=testing password=testing sslmode=disable")
	ResetTables()
}

func ResetTables() {
	db.DropTable(&FileArtifact{})
	db.DropTable(&ScreenShot{})
	db.DropTable(&Ticket{})

	db.CreateTable(&FileArtifact{})
	db.CreateTable(&ScreenShot{})
	db.CreateTable(&Ticket{})
}
