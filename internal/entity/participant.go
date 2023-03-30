package entity

type Participant struct {
	ID           int    `json:"id`
	Name         string `json:"name"`
	Wish         string `json:"wish"`
	Recipient_id int    `json:"recipient_id"`
}
