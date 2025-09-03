package repository

import (
	"context"
	"database/sql"

	"github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/domain"
)

type UsersRepository interface {
	Create(ctx context.Context, user domain.User) (int64, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}

type UserPostgresRepository struct {
	DB *sql.DB
}

func NewUsersRepository(db *sql.DB) UsersRepository {
	return &UserPostgresRepository{
		DB: db,
	}
}

func (ur *UserPostgresRepository) Create(ctx context.Context, user domain.User) (int64, error) {
	q := `INSERT INTO users (name, email, password) 
          VALUES ($1, $2, $3) 
          RETURNING id`

	var id int64
	err := ur.DB.QueryRowContext(ctx, q, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ur *UserPostgresRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	u := &domain.User{}
	query := `SELECT name, email, password FROM users WHERE email = $1`
	err := ur.DB.QueryRowContext(ctx, query, email).Scan(&u.Name, &u.Email, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}
