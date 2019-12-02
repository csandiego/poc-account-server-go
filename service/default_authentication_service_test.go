package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGivenUserCredentialWhenDefaultAuthenticationServiceAuthenticateThenHashPassword(t *testing.T) {
	dao := testUserCredentialDao{}
	hasher := testPasswordHasher{"HASH"}
	service := NewDefaultAuthenticationService(&dao, &hasher)
	service.Authenticate(credential)
	require.Equal(t, dao.authenticateCredential.Password, hasher.hash)
}
