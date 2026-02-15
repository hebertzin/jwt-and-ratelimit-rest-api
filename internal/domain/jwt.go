package domain

type JwtService interface {
	GenerateToken(email string, userId string) (string, error)
	VerifyToken(token string) error
}
