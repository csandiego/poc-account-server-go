package service

import (
	"github.com/csandiego/poc-account-server/data"
)

type UserRegistrationService interface {
	Validate(string) (bool, error)
	Registration(data.UserCredential) error
}
