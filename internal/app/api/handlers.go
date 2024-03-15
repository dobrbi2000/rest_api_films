package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dobrbi2000/rest_api_films/internal/app/api/middleware"
	"github.com/dobrbi2000/rest_api_films/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
)

//Доп структура для формирования сообщений

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeader(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (api *API) GetAllFilms(writer http.ResponseWriter, req *http.Request) {
	//инилизация хедера
	initHeader(writer)
	//логируем момент начала обработки запроса
	api.logger.Info("Get all Films GET /api/v1/films")

	films, err := api.storage.Film().SelectAll()
	if err != nil {
		api.logger.Info("Error while Films.SelectAll:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Database unavailible. Try again later",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(films)
}

func (api *API) PostFilms(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
	api.logger.Info("Post film POST /api/v1/films")

	var requestData struct {
		Film   models.Film     `json:"film"`
		Actors []*models.Actor `json:"actors"`
	}

	err := json.NewDecoder(req.Body).Decode(&requestData)
	fmt.Println("Request", req.Body)
	fmt.Println("Film:", requestData.Film)
	fmt.Println("Actors", requestData.Actors)
	defer req.Body.Close()

	if err != nil {
		api.logger.Info("Invalid JSON file recieved from client for film")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided JSON is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	film, actors := requestData.Film, requestData.Actors

	f, a, err := api.storage.Film().Create(&film, actors)
	fmt.Println("FilmCreate:", &film)
	fmt.Println("ActorsCreate", actors)
	if err != nil {
		api.logger.Info("Some error with creating new film:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some problem with the database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	data := struct {
		Film   models.Film
		Actors []*models.Actor
	}{
		Film:   *f,
		Actors: a,
	}

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(data)

}

func (api *API) GetFilmById(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
	api.logger.Info("Get Film by ID /api/v1/films/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Trouble while parsing {id} patameter:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "ID invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	film, ok, err := api.storage.Film().FindFilmById(id)
	if err != nil {
		api.logger.Info("Some trobles with acces to DB", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We don't have access to DB, sorry",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Cannot find film ID in DB")
		msg := Message{
			StatusCode: 404,
			Message:    "Film with this ID doesn't exist",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)

	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(film)

}
func (api *API) DeleteFilmById(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
	api.logger.Info("Delete film by ID")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Trouble while parsing {id} patameter:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "ID invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok, err := api.storage.Film().FindFilmById(id)
	if err != nil {
		api.logger.Info("Some trobles with acces to DB", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We don't have access to DB, sorry",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Cannot find article ID in DB")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with this ID doesn't exist",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)

	}
	_, err = api.storage.Film().DeleteById(id)
	if err != nil {
		api.logger.Info("Some trobles with delete from DB", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We don't have access to DB, sorry",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Film with ID %d successfully deleted", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)

}

func (api *API) PostActors(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
	api.logger.Info("Post actor POST /api/v1/actors")
	var actor models.Actor
	err := json.NewDecoder(req.Body).Decode(&actor)
	if err != nil {
		api.logger.Info("Invalid JSON file recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided JSON is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	a, err := api.storage.Actor().Create(&actor)
	if err != nil {
		api.logger.Info("Some error with creating new actor:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some problem with database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)

}

func (api *API) GetAllActors(writer http.ResponseWriter, req *http.Request) {
	//инилизация хедера
	initHeader(writer)
	//логируем момент начала обработки запроса
	api.logger.Info("Get all Films GET /api/v1/actors")

	actors, err := api.storage.Actor().SelectAll()
	if err != nil {
		api.logger.Info("Error while Actor.SelectAll:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Database unavailible. Try again later",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(actors)
}

func (api *API) PostUserRegister(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
	api.logger.Info("Post User Register POST /api/v1/user/register")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid JSON file recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided JSON is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//пытаемся найти пользователя с таким же логином в БД
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Some trobles with acces to DB (users)", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We don't have access to DB (users), sorry",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//проверка дубликата пользователя
	if ok {
		api.logger.Info("User with ID already exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User with ID already exists",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	userAdded, err := api.storage.User().Create(&user)
	if err != nil {
		api.logger.Info("Some trobles with acces to DB", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We don't have access to DB, sorry",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User {login:%s} successfully registered!", userAdded.Login),
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)

}

func (api *API) PostToAuth(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
	api.logger.Info("Post to Auth POST /api/v1/user/auth")
	var userFromJson models.User
	err := json.NewDecoder(req.Body).Decode(&userFromJson)
	//если невалидный json
	if err != nil {
		api.logger.Info("Invalid JSON file recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided JSON is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//проверка на существование юзера
	userInDb, ok, err := api.storage.User().FindByLogin(userFromJson.Login)
	if err != nil {
		api.logger.Info("Can't make user search in database:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles while accessing database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	//нет юзера в БД
	if !ok {
		api.logger.Info("User with ID doesn't exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User with ID doesn't exists exists. Try register first",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	//проверка пароля из запроса в БД
	if userInDb.Password != userFromJson.Password {
		api.logger.Info("Invalid credentials to auth")
		msg := Message{
			StatusCode: 404,
			Message:    "Your password is invalid",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)

	}

	//токен
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	claims["admin"] = true
	claims["name"] = userInDb

	tokenString, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		api.logger.Info("Can't claim jwt-token")
		msg := Message{
			StatusCode: 500,
			Message:    "We hava some troubles. Try later",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
	}
	msg := Message{
		StatusCode: 201,
		Message:    tokenString,
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)

}
