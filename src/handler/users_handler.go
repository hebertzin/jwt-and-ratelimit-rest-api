package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jwt-and-ratelimit-rest-api/src/domain"
	"github.com/jwt-and-ratelimit-rest-api/src/services"
)

type (
	UserHandler struct {
		UserService *services.UserService
		BaseHandler
	}

	createUserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.RespondWithError(w, r, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	u := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	id, err := handler.UserService.Create(r.Context(), u)
	if err != nil {
		handler.RespondWithError(w, r, err.Code, err.Error(), err.Message)
		return
	}

	handler.RespondWithSuccess(w, http.StatusCreated, "user created successfully", map[string]int64{"user_id": id})
}
