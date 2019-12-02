package service

type PasswordHasher interface {
	Hash(string) string
}
