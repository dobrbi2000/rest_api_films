package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dobrbi2000/rest_api_films/models"
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

func (api *API) GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	//инилизация хедера
	initHeader(writer)
	//логируем момент начала обработки запроса
	api.logger.Info("Get all Articles GET /api/v1/articles")
	//r := storage.NewArticleRepository(/**/)

	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		api.logger.Info("Error while Articles.SelectAll:", err)
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
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) PostArticle(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
	api.logger.Info("Post article POST /api/v1/articles")
	var article models.Article
	err := json.NewDecoder(req.Body).Decode(&article)
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
	a, err := api.storage.Article().Create(&article)
	if err != nil {
		api.logger.Info("Some error wit creating new article:", err)
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

func (api *API) GetArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
	api.logger.Info("Get Article by ID /api/v1/articles/{id}")
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
	article, ok, err := api.storage.Article().FindArticleById(id)
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
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)

}
func (api *API) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
	api.logger.Info("Delete by ID")
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
	_, ok, err := api.storage.Article().FindArticleById(id)
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
	_, err = api.storage.Article().DeleteById(id)
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
		Message:    fmt.Sprintf("Article with ID %d successfully deleted", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)

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
