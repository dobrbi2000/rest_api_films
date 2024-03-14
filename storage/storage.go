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
	db                *sql.DB
	userRepository    *UserRepository
	articleRepository *ArticleRepository
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

func (s *Storage) Article() *ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}
	s.articleRepository = &ArticleRepository{
		storage: s,
	}
	fmt.Println("initialize article storage:", s)
	return s.articleRepository
}
