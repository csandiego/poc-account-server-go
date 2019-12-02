package redigo

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestGivenNonExistingEmailWhenRedigoUserCredentialDaoEmailExistsThenReturnFalse(t *testing.T) {
	r := startRedis(t)
	defer r.Close()
	pool := createPool(r)
	defer pool.Close()
	dao := NewRedigoUserCredentialDao(pool)
	exists, err := dao.EmailExists(credential.Email)
	require.Nil(t, err)
	require.False(t, exists)
}

func TestGivenExistingEmailWhenRedigoUserCredentialDaoEmailExistsThenReturnTrue(t *testing.T) {
	r := startRedis(t)
	defer r.Close()
	pool := createPool(r)
	defer pool.Close()
	loadUserCredential(t, r)
	dao := NewRedigoUserCredentialDao(pool)
	exists, err := dao.EmailExists(credential.Email)
	require.Nil(t, err)
	require.True(t, exists)
}

func TestGivenValidUserCredentialWhenRedigoUserCredentialDaoCreateThenIncrementUserIdCounter(t *testing.T) {
	r := startRedis(t)
	defer r.Close()
	pool := createPool(r)
	defer pool.Close()
	dao := NewRedigoUserCredentialDao(pool)
	require.Nil(t, dao.Create(credential))
	id, err := r.Get(userIdCounterKey)
	require.Nil(t, err)
	require.Equal(t, "1", id)
}

func TestGivenValidUserCredentialWhenRedigoUserCredentialDaoCreateThenAddToStore(t *testing.T) {
	r := startRedis(t)
	defer r.Close()
	pool := createPool(r)
	defer pool.Close()
	dao := NewRedigoUserCredentialDao(pool)
	require.Nil(t, dao.Create(credential))
	require.True(t, r.Exists(fmt.Sprintf(userCredentialKeyFmt, credential.Email)))
}

func TestGivenValidUserCredentialWhenRedigoUserCredentialDaoAuthenticateThenReturnUserId(t *testing.T) {
	r := startRedis(t)
	defer r.Close()
	pool := createPool(r)
	defer pool.Close()
	loadUserCredential(t, r)
	dao := NewRedigoUserCredentialDao(pool)
	id, err := dao.Authenticate(credential)
	require.Nil(t, err)
	userId, err := strconv.ParseInt(r.HGet(fmt.Sprintf(userCredentialKeyFmt, credential.Email), userCredentialUserIdKey), 10, 64)
	require.Nil(t, err)
	require.Equal(t, userId, int64(id))
}

func TestGivenNonExistingUserCredentialWhenRedigoUserCredentialDaoAuthenticateThenReturnError(t *testing.T) {
	r := startRedis(t)
	defer r.Close()
	pool := createPool(r)
	defer pool.Close()
	dao := NewRedigoUserCredentialDao(pool)
	_, err := dao.Authenticate(credential)
	require.NotNil(t, err)
}

func TestGivenUserCredentialWithWrongPasswordWhenRedigoUserCredentialDaoAuthenticateThenReturnError(t *testing.T) {
	r := startRedis(t)
	defer r.Close()
	pool := createPool(r)
	defer pool.Close()
	loadUserCredential(t, r)
	fakeCredential := credential
	fakeCredential.Password = "INCORRECT"
	dao := NewRedigoUserCredentialDao(pool)
	_, err := dao.Authenticate(fakeCredential)
	require.Equal(t, ErrPasswordMismatch, err)
}
