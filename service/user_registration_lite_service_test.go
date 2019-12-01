package service

import (
	"errors"
	"github.com/csandiego/poc-account-server/data"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGivenNonExistingEmailWhenUserRegistrationLiteServiceValidateThenReturnTrue(t *testing.T) {
	dao := testUserCredentialDao{emailExists: false, emailExistsErr: nil}
	service := NewUserRegistrationLiteService(&dao)
	valid, err := service.Validate("someone@somewhere.com")
	require.Nil(t, err)
	require.True(t, valid)
}

func TestGivenExistingEmailWhenUserRegistrationLiteServiceValidateThenReturnFalse(t *testing.T) {
	dao := testUserCredentialDao{emailExists: true, emailExistsErr: nil}
	service := NewUserRegistrationLiteService(&dao)
	valid, err := service.Validate("someone@somewhere.com")
	require.Nil(t, err)
	require.False(t, valid)
}

func TestGivenDaoHasErrorsWhenUserRegistrationLiteServiceValidateThenReturnError(t *testing.T) {
	dao := testUserCredentialDao{emailExistsErr: errors.New("TEST")}
	service := NewUserRegistrationLiteService(&dao)
	_, err := service.Validate("someone@somewhere.com")
	require.NotNil(t, err)
}

func TestGivenValidUserCredentialWhenUserRegistrationLiteServiceRegisterThenNoError(t *testing.T) {
	dao := testUserCredentialDao{insertErr: nil}
	credential := data.UserCredential{Email: "someone@somewhere.com", Password: "password"}
	service := NewUserRegistrationLiteService(&dao)
	require.Nil(t, service.Register(credential))
}

func TestGivenDaoHasErrorsWhenUserRegistrationLiteServiceRegisterThenReturnError(t *testing.T) {
	dao := testUserCredentialDao{insertErr: errors.New("TEST")}
	credential := data.UserCredential{Email: "someone@somewhere.com", Password: "password"}
	service := NewUserRegistrationLiteService(&dao)
	require.NotNil(t, service.Register(credential))
}
