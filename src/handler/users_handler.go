package handler

import (
	"net/http"

	"github.com/jwt-and-ratelimit-rest-api/src/domain"
	"github.com/jwt-and-ratelimit-rest-api/src/services"
	"github.com/jwt-and-ratelimit-rest-api/src/utils"
)

type UserHandler struct {
	UserService  services.UserService
	HttpResponse utils.BaseHandler
}

type createUserRequest struct {
	Name     string
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (uh *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	w.Header().Set("Content-Type", "application/json")

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	result, err := uh.UserService.Create(r.Context(), user)
	if err != nil {
		uh.HttpResponse.RespondWithError(w, r, err.Code, err.Error(), err.Message)
		return
	}

	uh.HttpResponse.RespondWithSuccess(w, http.StatusCreated, "user created successfully", result)
}
