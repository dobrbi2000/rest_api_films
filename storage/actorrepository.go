package storage

import (
	"fmt"
	"log"

	"github.com/dobrbi2000/rest_api_films/models"
)

type ActorRepository struct {
	storage *Storage
}

var (
	tableFilms     string = "films"
	tableActors    string = "actors"
	tableFilmActor string = "filmactor"
)

func (ac *ActorRepository) Create(a *models.Actor) (*models.Actor, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, gender, birth_date) VALUES ($1, $2, $3) RETURNING actor_id", tableActors)
	if err := ac.storage.db.QueryRow(query, a.Name, a.Gender, a.BirthDate).Scan(&a.ActorID); err != nil {
		return nil, err
	}

	return a, nil
}

func (ac *ActorRepository) DeleteById(id int) (*models.Actor, error) {
	actor, ok, err := ac.FindActorById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE actor_id=$1", tableActors)
		_, err := ac.storage.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}

	return actor, nil
}

func (ac *ActorRepository) FindActorById(id int) (*models.Actor, bool, error) {
	actor, err := ac.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var actorFinded *models.Actor
	for _, a := range actor {
		if a.ActorID == id {
			actorFinded = a
			founded = true
			break
		}
	}

	return actorFinded, founded, nil
}

func (ac *ActorRepository) SelectAll() ([]*models.Actor, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableActors)
	rows, err := ac.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	actors := make([]*models.Actor, 0)
	for rows.Next() {
		a := models.Actor{}
		err := rows.Scan(&a.ActorID, &a.Name, &a.Gender, &a.BirthDate)
		if err != nil {
			log.Println(err)
			continue
		}
		actors = append(actors, &a)
	}

	return actors, nil
}
