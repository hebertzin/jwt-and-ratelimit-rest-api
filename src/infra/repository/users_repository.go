package repository

import (
	"context"
	"database/sql"

	"github.com/jwt-and-ratelimit-rest-api/src/domain"
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
	tx := ur.DB.QueryRowContext(ctx, q, user.Name, user.Email, user.Password)
	tx.Scan(&id)
	if tx != nil {
		return 0, nil
	}

	return id, nil
}

func (ur *UserPostgresRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	u := domain.User{}
	q := `SELECT * FROM users WHERE email = ?`
	tx := ur.DB.QueryRowContext(ctx, q, email)
	tx.Scan(&u.Name, &u.Email, &u.Password)
	return &domain.User{}, tx.Err()
}
