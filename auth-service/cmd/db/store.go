package db

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"github.com/sum-project/CookiApp/auth-service/cmd/repository"
	"github.com/sum-project/CookiApp/auth-service/cmd/repository/postgres"
	redisrepo "github.com/sum-project/CookiApp/auth-service/cmd/repository/redis"
)

type Store struct {
	UserRepository  repository.UserRepository
	TokenRepository repository.TokenRepository
}

func NewStore(conn *sql.DB, client *redis.Client) *Store {
	return &Store{
		UserRepository:  postgres.NewUserRepository(conn),
		TokenRepository: redisrepo.NewTokenRepository(client),
	}
}
