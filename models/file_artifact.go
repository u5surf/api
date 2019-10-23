package models

type FileArtifact struct {
	TicketId uint   `json:"ticketID" gorm:"primary_key" sql:"not null"`
	Filename string `json:"filename" gorm:"primary_key,size:255" sql:"not null"`
	Data     []byte
}
