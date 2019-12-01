package service

import (
	"github.com/csandiego/poc-account-server/dao"
	"github.com/csandiego/poc-account-server/data"
)

type UserRegistrationLiteService struct {
	dao dao.UserCredentialDao
}

func NewUserRegistrationLiteService(dao dao.UserCredentialDao) *UserRegistrationLiteService {
	return &UserRegistrationLiteService{dao}
}

func (service *UserRegistrationLiteService) Validate(email string) (bool, error) {
	exists, err := service.dao.EmailExists(email)
	return !exists, err
}

func (service *UserRegistrationLiteService) Register(credential data.UserCredential) error {
	return service.dao.Insert(credential)
}
