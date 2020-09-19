package service

import (
	"github.com/csandiego/poc-account-server/dao"
	"github.com/csandiego/poc-account-server/data"
)

type DefaultUserProfileService struct {
	dao dao.UserProfileDao
}

func NewDefaultUserProfileService(dao dao.UserProfileDao) *DefaultUserProfileService {
	return &DefaultUserProfileService{dao: dao}
}

func (service *DefaultUserProfileService) Get(userId int) (*data.UserProfile, error) {
	return service.dao.Get(userId)
}

func (service *DefaultUserProfileService) Update(profile data.UserProfile) error {
	return service.dao.Update(profile)
}
