package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/sum-project/CookiApp/auth/cmd/models"
	"github.com/sum-project/CookiApp/auth/cmd/repository"
	"time"
)

const dbTimeout = time.Second * 3

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{DB: db}
}

func (u *userRepository) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, password, role, created_at, updated_at, deleted_at 
				from users where email = $1 and deleted_at is null`

	var user models.User
	row := u.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) GetUserByID(id uuid.UUID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, password, role, created_at, updated_at, deleted_at 
				from users where id = $1 and deleted_at is null`

	var user models.User
	row := u.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) InsertUser(user models.User) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `insert into users (id, email, password, role) values ($1, $2, $3, $4) returning id`

	var newID uuid.UUID

	err := u.DB.QueryRowContext(ctx, stmt,
		user.ID,
		user.Email,
		user.Password,
		user.Role,
	).Scan(&newID)

	if err != nil {
		return [16]byte{}, err
	}

	return newID, nil
}
