package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGivenNonExistingEmailWhenDefaultUserRegistrationServiceValidateThenReturnTrue(t *testing.T) {
	dao := testUserCredentialDao{emailExists: false}
	hasher := testPasswordHasher{}
	service := NewDefaultUserRegistrationService(&dao, &hasher)
	valid, _ := service.Validate(credential.Email)
	require.True(t, valid)
}

func TestGivenExistingEmailWhenDefaultUserRegistrationServiceValidateThenReturnFalse(t *testing.T) {
	dao := testUserCredentialDao{emailExists: true}
	hasher := testPasswordHasher{}
	service := NewDefaultUserRegistrationService(&dao, &hasher)
	valid, _ := service.Validate(credential.Email)
	require.False(t, valid)
}

func TestGivenValidUserCredentialWhenDefaultUserRegistrationServiceRegisterThenNoError(t *testing.T) {
	dao := testUserCredentialDao{}
	hasher := testPasswordHasher{"HASH"}
	service := NewDefaultUserRegistrationService(&dao, &hasher)
	service.Register(credential)
	require.Equal(t, hasher.hash, dao.createCredential.Password)
}
