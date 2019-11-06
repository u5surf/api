package models

import (
	"testing"
)

func Init() {
	ConnectToTesting()
	PopulateDB()
}

func TestGetAssociatedArtifacts(t *testing.T) {
	Init()
	artifacts, err := GetAssociatedArtifacts(1)
	if err != nil {
		t.Error("Error when getting artifacts associated with a ticket ID")
	}
	if len(*artifacts) != 2 ||
		(*artifacts)[0].Filename != "test1" || (*artifacts)[1].Filename != "test2" {
		t.Error("GetAssociatedArtifacts returns wrong entries")
	}

	artifacts, err = GetAssociatedArtifacts(2)
	if err != nil {
		t.Error("Error when getting artifacts associated with a ticket ID")
	}
	if len(*artifacts) != 1 ||
		((*artifacts)[0].Filename != "test3") {
		t.Error("GetAssociatedArtifacts returns wrong entries")
	}

	artifacts, err = GetAssociatedArtifacts(5)
	if err != nil {
		t.Error("Error when getting artifacts associated with a ticket ID")
	}
	if len(*artifacts) != 0 {
		t.Error("GetAssociatedArtifacts returns wrong entries")
	}
}

func TestGetArtifact(t *testing.T) {
	Init()
	artifact, err := GetArtifact(1, "test1")
	if err != nil {
		t.Error("Error when getting artifacts associated with a ticket ID")
	}
	if artifact == nil {
		t.Error("GetArtifact cannot find entiry")
	}

	artifact, err = GetArtifact(2, "test1")
	if err == nil {
		t.Error("GetArtifact did not report error when find an artifact which does not exist")
	}
}

func PopulateDB() {
	artifacts := []FileArtifact{
		FileArtifact{TicketId: 1, Filename: "test1", Data: []byte("TEST1")},
		FileArtifact{TicketId: 1, Filename: "test2", Data: []byte("TEST2")},
		FileArtifact{TicketId: 2, Filename: "test3", Data: []byte("TEST3")},
		FileArtifact{TicketId: 3, Filename: "test4", Data: []byte("TEST3")},
	}
	for _, a := range artifacts {
		db.Create(&a)
	}
}
