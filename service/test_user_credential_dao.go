package service

import (
	"github.com/csandiego/poc-account-server/data"
)

type testUserCredentialDao struct {
	emailExists    bool
	emailExistsErr error
	insertErr      error
}

func (dao *testUserCredentialDao) EmailExists(string) (bool, error) {
	return dao.emailExists, dao.emailExistsErr
}

func (dao *testUserCredentialDao) Insert(data.UserCredential) error {
	return dao.insertErr
}
