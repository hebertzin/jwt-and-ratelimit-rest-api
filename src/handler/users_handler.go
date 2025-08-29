package handler

import (
	"net/http"

	"github.com/jwt-and-ratelimit-rest-api/src/domain"
	"github.com/jwt-and-ratelimit-rest-api/src/services"
)

type UserHandler struct {
	UserService *services.UserService
	BaseHandler
}

type createUserRequest struct {
	Name     string
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserHanlder(userService *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req createUserRequest

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	result, err := handler.UserService.Create(r.Context(), user)
	if err != nil {
		handler.RespondWithError(w, r, err.Code, err.Error(), err.Message)
		return
	}

	handler.RespondWithSuccess(w, http.StatusCreated, "user created successfully", result)
}
