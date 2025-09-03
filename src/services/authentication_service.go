package services

import (
	"context"

	"github.com/jwt-and-ratelimit-rest-api/src/infra/repository"
	"github.com/jwt-and-ratelimit-rest-api/src/infra/security"
	"github.com/jwt-and-ratelimit-rest-api/src/utils"
)

type AuthenticationService struct {
	repo repository.UsersRepository
	hash security.PasswordHasher
}

func NewAuthenticationService(repo repository.UsersRepository, hash security.PasswordHasher) *AuthenticationService {
	return &AuthenticationService{repo: repo, hash: hash}
}

func (s *AuthenticationService) AuthenticateUser(ctx context.Context, email, password string) (string, *utils.Exception) {
	u, _ := s.repo.FindByEmail(ctx, email)
	if u == nil {
		return "", utils.BadRequest(utils.WithMessage("user not found"))
	}

	err := s.hash.Compare(password, u.Password)
	if err != nil {
		return "", utils.BadRequest(utils.WithMessage("invalid credentials"))
	}

	token, err := security.CreateToken(email)
	if err != nil {
		return "", utils.BadRequest(utils.WithMessage(err.Error()))
	}

	return token, nil
}
