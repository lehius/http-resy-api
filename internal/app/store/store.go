package store

import (
	"database/sql"

	"github.com/lehius/http-resy-api/internal/config"
	_ "github.com/lib/pq" // ...
)

const (
	scheme = "postgres"
)

// Store ...
type Store struct {
	config         *config.Store
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(config *config.Store) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open(scheme, s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// User ...
func (s *Store) User() *UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}

	return s.userRepository
}

// store.User().Create()
