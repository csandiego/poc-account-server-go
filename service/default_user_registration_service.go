package service

import (
	"github.com/csandiego/poc-account-server/dao"
	"github.com/csandiego/poc-account-server/data"
)

type DefaultUserRegistrationService struct {
	dao dao.UserCredentialDao
	hasher PasswordHasher
}

func NewDefaultUserRegistrationService(dao dao.UserCredentialDao, hasher PasswordHasher) *DefaultUserRegistrationService {
	return &DefaultUserRegistrationService{dao, hasher}
}

func (service *DefaultUserRegistrationService) Validate(email string) (bool, error) {
	exists, err := service.dao.EmailExists(email)
	return !exists, err
}

func (service *DefaultUserRegistrationService) Register(credential data.UserCredential) error {
	credential.Password = service.hasher.Hash(credential.Password)
	return service.dao.Create(credential)
}
