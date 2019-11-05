package models

type FileArtifact struct {
	TicketId uint   `json:"ticketID" gorm:"primary_key"`
	Filename string `json:"filename" gorm:"primary_key"`
	Data     []byte `json:"-"`
}

func GetAssociatedArtifacts(TicketID uint) (*[]FileArtifact, error) {
	var fileArtifacts []FileArtifact

	err := db.Where("ticket_id = ?", TicketID).Find(&fileArtifacts, TicketID).Error

	return &fileArtifacts, err
}
