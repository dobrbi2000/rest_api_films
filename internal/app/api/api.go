package api

import (
	"net/http"

	"github.com/dobrbi2000/web/rest_api_films/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// base api server
type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

// api constructor
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {
	//configure logger
	if err := api.configireLoggerField(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port:", api.config.BindAddr)

	api.configireRouterField()

	if err := api.configireStorageField(); err != nil {
		return err
	}

	return http.ListenAndServe(api.config.BindAddr, api.router)
}
