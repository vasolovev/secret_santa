package entity

type Group struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Participants []Participant `json:"participants"`
}
