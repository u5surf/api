package models

type FileArtifact struct {
	TicketId uint   `json:"ticketID" gorm:"primary_key"`
	Filename string `json:"filename" gorm:"primary_key"`
	Data     []byte
}
