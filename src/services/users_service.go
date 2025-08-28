package services

import (
	"context"

	"github.com/jwt-and-ratelimit-rest-api/src/domain"
	"github.com/jwt-and-ratelimit-rest-api/src/infra/repository"
	"github.com/jwt-and-ratelimit-rest-api/src/utils"
	"github.com/jwt-and-ratelimit-rest-api/src/utils/validation"
)

type UserService struct {
	repo              repository.UsersRepository
	payloadValidation validation.PayloadValidation
}

type CreateUserPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

func NewUserService(repo repository.UsersRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, user domain.User) (int64, *utils.Exception) {
	var payload CreateUserPayload
	if err := s.payloadValidation.ValidateStruct(payload); err != nil {
		return 0, utils.BadRequest(utils.WithMessage("error validating payload"))
	}

	id, err := s.repo.Create(ctx, user)
	if err != nil {
		return id, utils.Unexpected(utils.WithMessage("error creating user"))
	}

	return id, nil
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (*domain.User, *utils.Exception) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, utils.Unexpected(utils.WithMessage("error finding user by email"))
	}

	return user, nil
}
