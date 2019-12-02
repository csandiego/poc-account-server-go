package service

import (
	"github.com/csandiego/poc-account-server/dao"
	"github.com/csandiego/poc-account-server/data"
)

type DefaultAuthenticationService struct {
	dao dao.UserCredentialDao
	hasher PasswordHasher
}

func NewDefaultAuthenticationService(dao dao.UserCredentialDao, hasher PasswordHasher) *DefaultAuthenticationService {
	return &DefaultAuthenticationService{dao, hasher}
}

func (service *DefaultAuthenticationService) Authenticate(credential data.UserCredential) (int, error) {
	credential.Password = service.hasher.Hash(credential.Password)
	return service.dao.Authenticate(credential)
}
