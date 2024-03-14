package api

import (
	"net/http"

	"github.com/dobrbi2000/rest_api_films/internal/app/api/middleware"
	"github.com/dobrbi2000/rest_api_films/storage"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
)

// конфигурация для API инстанса (Logger)
func (a *API) configireLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil

}

// конфиг для роутера
func (a *API) configireRouterField() {

	a.router.HandleFunc(prefix+"/films", a.GetAllFilms).Methods("GET")
	//a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	a.router.Handle(prefix+"/film/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.GetFilmById),
	)).Methods("GET")
	//a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.Handle(prefix+"/articles/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.DeleteFilmById),
	)).Methods("DELETE")
	a.router.HandleFunc(prefix+"/films", a.PostFilms).Methods("POST")

	a.router.HandleFunc(prefix+"/actors", a.PostActors).Methods("POST")
	a.router.HandleFunc(prefix+"/actors", a.GetAllActors).Methods("GET")
	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")
	//конфиг для auth
	a.router.HandleFunc(prefix+"/user/auth", a.PostToAuth).Methods("POST")

}

// конфиг БД
func (a *API) configireStorageField() error {
	storage := storage.New(a.config.Storage)
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil

}
