package services

import (
	"context"

	"github.com/hebertzin/jwt-and-ratelimit-rest-api/internal/domain"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/internal/infra/repository"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/internal/infra/security"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/internal/utils"
)

type AuthenticationService struct {
	repo repository.UsersRepository
	hash security.PasswordHasher
	jwt  domain.JwtService
}

func NewAuthenticationService(repo repository.UsersRepository, hash security.PasswordHasher, jwt domain.JwtService) *AuthenticationService {
	return &AuthenticationService{
		repo: repo,
		hash: hash,
		jwt:  jwt,
	}
}

func (s *AuthenticationService) AuthenticateUser(ctx context.Context, email, password string) (string, *utils.Exception) {
	u, _ := s.repo.FindByEmail(ctx, email)
	if u == nil {
		return "", utils.BadRequest(utils.WithMessage("user not found"))
	}

	err := s.hash.Compare(u.Password, password)
	if err != nil {
		return "", utils.BadRequest(utils.WithMessage("invalid credentials"))
	}

	token, err := s.jwt.GenerateToken(email, u.ID)
	if err != nil {
		return "", utils.BadRequest(utils.WithMessage(err.Error()))
	}

	return token, nil
}
