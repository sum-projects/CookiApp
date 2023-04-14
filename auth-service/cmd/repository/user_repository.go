package repository

import (
	"github.com/google/uuid"
	"github.com/sum-project/CookiApp/auth-service/cmd/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	UpdateUser(u *models.User) error
	InsertUser(u models.User) (uuid.UUID, error)
}
