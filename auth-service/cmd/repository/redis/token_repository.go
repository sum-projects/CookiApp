package redis

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/sum-project/CookiApp/auth-service/cmd/repository"
	"time"
)

type tokenRepository struct {
	client *redis.Client
}

func NewTokenRepository(client *redis.Client) repository.TokenRepository {
	return &tokenRepository{
		client: client,
	}
}

func (r tokenRepository) GenerateToken(id string) (string, error) {
	token := generateToken()

	if err := r.client.Set(context.Background(), token, id, time.Hour*24).Err(); err != nil {
		return "", err
	}

	return token, nil
}

func (r tokenRepository) ValidToken(token string) (string, error) {
	id, err := r.client.Get(context.Background(), token).Result()
	if err != nil && id != "" {
		return "", err
	}

	return id, nil
}

func generateToken() string {
	return uuid.New().String()
}
