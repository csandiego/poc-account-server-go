package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGivenNonExistingEmailWhenUserRegistrationLiteServiceValidateThenReturnTrue(t *testing.T) {
	dao := testUserCredentialDao{emailExists: false}
	hasher := testPasswordHasher{}
	service := NewUserRegistrationLiteService(&dao, &hasher)
	valid, _ := service.Validate(credential.Email)
	require.True(t, valid)
}

func TestGivenExistingEmailWhenUserRegistrationLiteServiceValidateThenReturnFalse(t *testing.T) {
	dao := testUserCredentialDao{emailExists: true}
	hasher := testPasswordHasher{}
	service := NewUserRegistrationLiteService(&dao, &hasher)
	valid, _ := service.Validate(credential.Email)
	require.False(t, valid)
}

func TestGivenValidUserCredentialWhenUserRegistrationLiteServiceRegisterThenNoError(t *testing.T) {
	dao := testUserCredentialDao{}
	hasher := testPasswordHasher{"HASH"}
	service := NewUserRegistrationLiteService(&dao, &hasher)
	service.Register(credential)
	require.Equal(t, dao.createCredential.Password, hasher.hash)
}
