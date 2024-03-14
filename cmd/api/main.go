package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/dobrbi2000/rest_api_films/internal/app/api"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	flag.Parse()
	log.Println("It works")

	//server init
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config) //десереализация toml файла

	if err != nil {
		log.Println("Can't find configs files", err)
	}
	server := api.New(config)

	log.Fatal(server.Start())

}
