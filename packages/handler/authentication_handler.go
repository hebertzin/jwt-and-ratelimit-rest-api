package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jwt-and-ratelimit-rest-api/packages/services"
)

type (
	AuthenticationHandler struct {
		authenticationService *services.AuthenticationService
		BaseHandler
	}

	authenticateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func NewAuthenticatorHandler(userService *services.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{
		authenticationService: userService,
	}
}

// Authenticate godoc
// @Summary      Authenticate user
// @Description  Authenticates a user with email and password and returns a JWT token
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Param        request body authenticateUserRequest true "User credentials"
// @Success      200 {object} HttpResponse "JWT token generated successfully"
// @Failure      400 {object} ErrorResponse "Invalid request body"
// @Failure      401 {object} ErrorResponse "Invalid credentials"
// @Router       /authentication/login [post]
func (handler *AuthenticationHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var req authenticateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.RespondWithError(w, r, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	t, err := handler.authenticationService.AuthenticateUser(r.Context(), req.Email, req.Password)
	if err != nil {
		handler.RespondWithError(w, r, err.Code, err.Error(), err.Message)
		return
	}

	handler.RespondWithSuccess(w, http.StatusOK, "authentication successfully", map[string]string{"token": t})
}
