package db

import (
	"database/sql"
	"github.com/sum-project/CookiApp/auth/cmd/repository"
	"github.com/sum-project/CookiApp/auth/cmd/repository/postgres"
)

type Store struct {
	UserRepository repository.UserRepository
}

func NewStore(conn *sql.DB) *Store {
	return &Store{
		UserRepository: postgres.NewUserRepository(conn),
	}
}
