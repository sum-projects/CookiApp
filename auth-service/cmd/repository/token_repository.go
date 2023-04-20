package repository

type TokenRepository interface {
	GenerateToken(string) (string, error)
	ValidToken(string) (string, error)
}
