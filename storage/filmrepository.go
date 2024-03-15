package storage

import (
	"fmt"
	"log"

	"github.com/dobrbi2000/rest_api_films/models"
)

type FilmRepository struct {
	storage *Storage
}

func (fi *FilmRepository) Create(film *models.Film, actors []*models.Actor) (*models.Film, []*models.Actor, error) {
	//создание фильма
	filmQuery := fmt.Sprintf("INSERT INTO %s (title, description, year, rating) VALUES ($1, $2, $3, $4) RETURNING film_id", tableFilms)
	err := fi.storage.db.QueryRow(filmQuery, film.Title, film.Description, film.Year, film.Rating).Scan(&film.FilmID)
	fmt.Println("filmQuery", filmQuery)
	fmt.Println("filmError", err)
	if err != nil {
		return nil, nil, err
	}
	//создание актеров
	for _, actor := range actors {
		actorQuery := fmt.Sprintf("INSERT INTO %s (name, gender, birth_date) VALUES ($1, $2, $3) RETURNING actor_id", tableActors)
		fmt.Println("actorError", err)
		fmt.Println("actorQuery", actorQuery)
		err := fi.storage.db.QueryRow(actorQuery, actor.Name, actor.Gender, actor.BirthDate).Scan(&actor.ActorID)
		if err != nil {
			return nil, nil, err
		}
		filmActorQuery := fmt.Sprintf("INSERT INTO %s (film_id, actor_id) VALUES ($1, $2)", tableFilmActor)
		fmt.Println("filmError", err)
		fmt.Println("filmActorQuery", filmActorQuery)
		_, err = fi.storage.db.Exec(filmActorQuery, film.FilmID, actor.ActorID)
		if err != nil {
			return nil, nil, err
		}
	}

	return film, actors, nil
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
		if f.FilmID == id {
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
		err := rows.Scan(&f.FilmID, &f.Title, &f.Description, &f.Year, &f.Rating)
		if err != nil {
			log.Println(err)
			continue
		}
		films = append(films, &f)
	}

	return films, nil
}
