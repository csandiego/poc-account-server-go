package service

import (
	"github.com/csandiego/poc-account-server/data"
)

type testUserCredentialDao struct {
	emailExists    bool
	emailExistsErr error
	createErr      error
}

func (dao *testUserCredentialDao) EmailExists(string) (bool, error) {
	return dao.emailExists, dao.emailExistsErr
}

func (dao *testUserCredentialDao) Create(data.UserCredential) error {
	return dao.createErr
}
