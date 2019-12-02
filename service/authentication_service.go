package service

import (
	"github.com/csandiego/poc-account-server/dao"
	"github.com/csandiego/poc-account-server/data"
)

type AuthenticationService struct {
	dao dao.UserCredentialDao
	hasher PasswordHasher
}

func NewAuthenticationService(dao dao.UserCredentialDao, hasher PasswordHasher) *AuthenticationService {
	return &AuthenticationService{dao, hasher}
}

func (service *AuthenticationService) Authenticate(credential data.UserCredential) (int, error) {
	credential.Password = service.hasher.Hash(credential.Password)
	return service.dao.Authenticate(credential)
}
