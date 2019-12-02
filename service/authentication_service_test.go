package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGivenUserCredentialWhenAuthenticationServiceAuthenticateThenHashPassword(t *testing.T) {
	dao := testUserCredentialDao{}
	hasher := testPasswordHasher{"HASH"}
	service := NewAuthenticationService(&dao, &hasher)
	service.Authenticate(credential)
	require.Equal(t, dao.authenticateCredential.Password, hasher.hash)
}
