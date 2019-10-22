package models

import "github.com/jinzhu/gorm"

type Ticket struct {
	gorm.Model
	ID         int        `json:"id",gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name       string     `json:"name",gorm:"size:255"`
	URL        string     `json:"url"gorm:"size:4096"`
	Screenshot ScreenShot `json:"screenshot"`
}
