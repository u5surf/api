package models

type Ticket struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	URL        string     `json:"url"`
	Screenshot ScreenShot `json:"screenshot"`
}
