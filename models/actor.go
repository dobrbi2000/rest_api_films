package models

type Actor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	FilmIDs   []int  `json:"film_ids"`
}
