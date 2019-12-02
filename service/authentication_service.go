package service

import (
	"github.com/csandiego/poc-account-server/data"
)

type AuthenticationService interface {
	Authenticate(data.UserCredential) (int, error)
}
