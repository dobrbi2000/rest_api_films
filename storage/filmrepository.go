package storage

import (
	"fmt"
	"log"

	"github.com/dobrbi2000/rest_api_films/models"
)

type FilmRepository struct {
	storage *Storage
}

func (fi *FilmRepository) Create(a *models.Film) (*models.Film, error) {
	filmQuery := fmt.Sprintf("INSERT INTO %s (title, description, year, rating, actors ) VALUES ($1, $2, $3, $4) RETURNING id", tableFilms)
	if err := fi.storage.db.QueryRow(filmQuery, a.Title, a.Description, a.Year, a.Rating).Scan(&a.ID); err != nil {
		return nil, err
	}
	//Create actors
	for _, actor := range a.Actors {
		actorQuery := fmt.Sprintf("INSERT INTO %s (name, gender, birth_date) VALUES ($1, $2, $3) RETURNING id", tableActors)
		if err := fi.storage.db.QueryRow(actorQuery, actor.Name, actor.Gender, actor.BirthDate).Scan(&actor.ID); err != nil {
			return nil, err
		}
	}

	return a, nil
}

func (fi *FilmRepository) DeleteById(id int) (*models.Film, error) {
	film, ok, err := fi.FindFilmById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableFilms)
		_, err := fi.storage.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}

	return film, nil
}

func (fi *FilmRepository) FindFilmById(id int) (*models.Film, bool, error) {
	film, err := fi.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var filmFinded *models.Film
	for _, f := range film {
		if f.ID == id {
			filmFinded = f
			founded = true
			break
		}
	}

	return filmFinded, founded, nil
}

func (fi *FilmRepository) SelectAll() ([]*models.Film, error) {
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY Rating DESC", tableFilms)
	rows, err := fi.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	films := make([]*models.Film, 0)
	for rows.Next() {
		f := models.Film{}
		err := rows.Scan(&f.ID, &f.Title, &f.Description, &f.Year, &f.Rating, &f.Actors)
		if err != nil {
			log.Println(err)
			continue
		}
		films = append(films, &f)
	}

	return films, nil
}
