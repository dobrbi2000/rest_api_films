package models

type Film struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Year        string  `json:"year"`
	Rating      int     `json:"rating"`
	Actors      []Actor `json:"actors"`
}
