package dao

import (
	"github.com/csandiego/poc-account-server/data"
)

type UserProfileDao interface {
	Create(data.UserProfile) error
	Get(int) (*data.UserProfile, error)
	Update(data.UserProfile) error
}
