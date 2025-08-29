package services

import (
	"context"

	"github.com/jwt-and-ratelimit-rest-api/src/domain"
	"github.com/jwt-and-ratelimit-rest-api/src/infra/repository"
	"github.com/jwt-and-ratelimit-rest-api/src/infra/security"
	"github.com/jwt-and-ratelimit-rest-api/src/utils"
	"github.com/jwt-and-ratelimit-rest-api/src/utils/validation"
)

type UserService struct {
	repo              repository.UsersRepository
	payloadValidation validation.PayloadValidate
	hasher            security.PasswordHasher
}

type CreateUserPayload struct {
	Name     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

func NewUserService(
	repo repository.UsersRepository,
	hash security.PasswordHasher,
	payloadValidation validation.PayloadValidate,
) *UserService {
	return &UserService{
		repo:              repo,
		hasher:            hash,
		payloadValidation: payloadValidation,
	}
}

func (s *UserService) Create(ctx context.Context, user domain.User) (int64, *utils.Exception) {
	u, _ := s.repo.FindByEmail(ctx, user.Email)
	if u != nil {
		return 0, utils.Confilct(utils.WithMessage("user already exists"))
	}

	payload := CreateUserPayload{
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}

	if err := s.payloadValidation.ValidateStruct(payload); err != nil {
		return 0, utils.BadRequest(utils.WithMessage("error validating payload"))
	}

	hash, err := s.hasher.Hash(user.Password)
	if err != nil {
		return 0, utils.BadRequest(utils.WithMessage(err.Error()))
	}

	user.Password = hash

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
