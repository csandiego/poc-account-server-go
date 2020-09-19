package service

import (
	"github.com/csandiego/poc-account-server/data"
)

type UserProfileService interface {
	Get(int) (*data.UserProfile, error)
	Update(data.UserProfile) error
}
