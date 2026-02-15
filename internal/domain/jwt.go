package domain

type JwtService interface {
	GenerateToken(email string) (string, error)
	VerifyToken(token string) error
}
