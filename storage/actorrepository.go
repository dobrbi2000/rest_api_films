package storage

import (
	"fmt"

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
	query := fmt.Sprintf("INSERT INTO %s (name, gender, birth_date) VALUES ($1, $2, $3) RETURNING id", tableActors)
	if err := ac.storage.db.QueryRow(query, a.Name, a.Gender, a.BirthDate).Scan(&a.ActorID); err != nil {
		return nil, err
	}

	return a, nil
}

// func (ac *ActorRepository) DeleteById(id int) (*models.Actor, error) {
// 	actor, ok, err := ac.FindActorById(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if ok {
// 		query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableActors)
// 		_, err := ac.storage.db.Exec(query, id)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return actor, nil
// }

// func (ac *ActorRepository) FindActorById(id int) (*models.Actor, bool, error) {
// 	actor, err := ac.SelectAll()
// 	var founded bool
// 	if err != nil {
// 		return nil, founded, err
// 	}
// 	var actorFinded *models.Actor
// 	for _, a := range actor {
// 		if a.ID == id {
// 			actorFinded = a
// 			founded = true
// 			break
// 		}
// 	}

// 	return actorFinded, founded, nil
// }

// func (ac *ActorRepository) SelectAll() ([]*models.Actor, error) {
// 	query := fmt.Sprintf("SELECT actor.id, actor.name, actor.gender, actor.birth_date, array_agg(films.id) as film_ids FROM %s AS actors LEFT JOIN films_actors ON actors.id = films_actors.actor_id LEFT JOIN %s AS films ON films.id = films_actors.film_id GROUP BY actors.id", tableActors, tableFilms)
// 	rows, err := ac.storage.db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	actorsMap := make(map[int]*models.Actor)
// 	for rows.Next() {
// 		a := models.Actor{}s
// 		filmIDs := []int{}
// 		err := rows.Scan(&a.ActorID, &a.Name, &a.Gender, &a.BirthDate, pq.Array(&filmIDs))
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}
// 		if actor, ok := actorsMap[a.ActorID]; ok {
// 			actor. = filmIDs
// 		} else {
// 			a.FilmIDs = filmIDs
// 			actorsMap[a.ID] = &a
// 		}
// 	}
// 	actors := make([]*models.Actor, 0)
// 	for _, actor := range actorsMap {
// 		actors = append(actors, actor)
// 	}

// 	return actors, nil
//}
