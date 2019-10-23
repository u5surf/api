package models

import (
	"github.com/jinzhu/gorm"
)

type Ticket struct {
	gorm.Model
	Name       string       `json:"name",gorm:"size:255"`
	URL        string       `json:"url",gorm:"size:4096"`
	ScreenShot []ScreenShot `json:"screenshots"`
	Processed  bool         `json:"processed"`
}

func CreateTicket(ticket *Ticket) error {
	return db.Create(ticket).Error
}

func GetTicketById(ID uint) (*Ticket, error) {
	var ticket Ticket

	// Preload line fetches the screenshot table and joins automatically:
	err := db.Preload("ScreenShot").First(&ticket, ID).Error

	return &ticket, err
}
