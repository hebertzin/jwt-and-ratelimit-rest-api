package services

import (
	"context"

	"github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/domain"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/infra/repository"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/infra/security"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/utils"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/utils/validation"
)

type UserService struct {
	repo              repository.UsersRepository
	payloadValidation validation.PayloadValidate
	hasher            security.PasswordHasher
}

type CreateUserPayload struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	IsActive bool
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
	existing, _ := s.repo.FindByEmail(ctx, user.Email)
	if existing != nil {
		return 0, utils.Confilct(utils.WithMessage("user already exists"))
	}

	payload := CreateUserPayload{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		IsActive: user.IsActive,
	}

	if err := s.payloadValidation.ValidateStruct(payload); err != nil {
		return 0, utils.BadRequest(utils.WithMessage("error validating payload"))
	}

	hashed, err := s.hasher.Hash(user.Password)
	if err != nil {
		return 0, utils.BadRequest(utils.WithMessage(err.Error()))
	}

	user.Password = hashed
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
