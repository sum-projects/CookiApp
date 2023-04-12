package db

import "database/sql"

type Store struct {
	DB *sql.DB
}

func NewStore(conn *sql.DB) *Store {
	return &Store{DB: conn}
}
