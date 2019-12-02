package service

import (
	"github.com/csandiego/poc-account-server/dao"
	"github.com/csandiego/poc-account-server/data"
)

type UserRegistrationLiteService struct {
	dao dao.UserCredentialDao
	hasher PasswordHasher
}

func NewUserRegistrationLiteService(dao dao.UserCredentialDao, hasher PasswordHasher) *UserRegistrationLiteService {
	return &UserRegistrationLiteService{dao, hasher}
}

func (service *UserRegistrationLiteService) Validate(email string) (bool, error) {
	exists, err := service.dao.EmailExists(email)
	return !exists, err
}

func (service *UserRegistrationLiteService) Register(credential data.UserCredential) error {
	credential.Password = service.hasher.Hash(credential.Password)
	return service.dao.Create(credential)
}
