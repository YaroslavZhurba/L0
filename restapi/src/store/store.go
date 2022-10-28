package store

import (
	"database/sql"
	"restapi/src/model"

	_ "github.com/lib/pq" // ...
)

// Store ...
 type Store struct {
 	config          *Config
 	db              *sql.DB
 	orderRepository *OrderRepository
 }

 // New ...
 func New(config *Config) *Store {
 	return &Store{
 		config: config}
 }

 // Open ...
 func (s *Store) Open() error {
 	db, err := sql.Open("postgres", s.config.DatabaseURL)
 	if err != nil {
 		return err
 	}

 	if err := db.Ping(); err != nil {
 		return err
 	}

 	s.db = db

 	return nil
 }

 // Close ...
 func (s *Store) Close() {
 	s.db.Close()
 }

 // Order ...
 func (s *Store) Order() *OrderRepository {
 	if s.orderRepository != nil {
 		return s.orderRepository
 	}

 	s.orderRepository = &OrderRepository{
 		store: s,
		cash: make(map[int]*model.Order),
 	}

 	return s.orderRepository
 }