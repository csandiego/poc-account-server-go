package service

import (
	"github.com/csandiego/poc-account-server/data"
)

var credential = data.UserCredential{Email: "someone@somewhere.com", Password: "password"}

type testPasswordHasher struct {
	hash string
}

func (hasher *testPasswordHasher) Hash(password string) string {
	return hasher.hash
}

type testUserCredentialDao struct {
	emailExists            bool
	emailExistsErr         error
	createCredential       data.UserCredential
	createErr              error
	authenticateCredential data.UserCredential
	authenticateUserId     int
	authenticateErr        error
}

func (dao *testUserCredentialDao) EmailExists(string) (bool, error) {
	return dao.emailExists, dao.emailExistsErr
}

func (dao *testUserCredentialDao) Create(credential data.UserCredential) error {
	dao.createCredential = credential
	return dao.createErr
}

func (dao *testUserCredentialDao) Authenticate(credential data.UserCredential) (int, error) {
	dao.authenticateCredential = credential
	return dao.authenticateUserId, dao.authenticateErr
}
