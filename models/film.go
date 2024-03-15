package models

type Film struct {
	FilmID      int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Rating      int    `json:"rating"`
}
