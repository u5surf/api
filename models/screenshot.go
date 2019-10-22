package models

import "github.com/jinzhu/gorm"

type ScreenShot struct {
	gorm.Model
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Filename string `json:"filename"`
}
