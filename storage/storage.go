package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	config *Config
	//database filedescriptor
	db              *sql.DB
	userRepository  *UserRepository
	actorRepository *ActorRepository
	filmRepository  *FilmRepository
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// open connection
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database connection created successfully!")
	return nil

}

func (storage *Storage) Close() {
	storage.db.Close()

}

func (s *Storage) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		storage: s,
	}
	return s.userRepository
}

func (s *Storage) Actor() *ActorRepository {
	if s.actorRepository != nil {
		return s.actorRepository
	}
	s.actorRepository = &ActorRepository{
		storage: s,
	}
	fmt.Println("initialize article storage:", s)
	return s.actorRepository
}

func (s *Storage) Film() *FilmRepository {
	if s.filmRepository != nil {
		return s.filmRepository
	}
	s.filmRepository = &FilmRepository{
		storage: s,
	}
	fmt.Println("initialize article storage:", s)
	return s.filmRepository
}
