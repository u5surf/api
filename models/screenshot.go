package models

type ScreenShot struct {
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Filename string `json:"filename"`
}
