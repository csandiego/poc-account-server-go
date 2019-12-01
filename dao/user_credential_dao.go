package dao

import (
	"github.com/csandiego/poc-account-server/data"
)

type UserCredentialDao interface {
	EmailExists(string) (bool, error)
	Create(data.UserCredential) error
}
