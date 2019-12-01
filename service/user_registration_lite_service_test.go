package service

import (
	"errors"
	"github.com/alicebob/miniredis"
	"github.com/csandiego/poc-account-server/dao"
	"github.com/csandiego/poc-account-server/data"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/require"
	"testing"
)

var credential = data.UserCredential{Email: "someone@somewhere.com", Password: "password"}

func TestGivenNonExistingEmailWhenUserRegistrationLiteServiceValidateThenReturnTrue(t *testing.T) {
	dao := testUserCredentialDao{emailExists: false, emailExistsErr: nil}
	service := NewUserRegistrationLiteService(&dao)
	valid, err := service.Validate(credential.Email)
	require.Nil(t, err)
	require.True(t, valid)
}

func TestGivenExistingEmailWhenUserRegistrationLiteServiceValidateThenReturnFalse(t *testing.T) {
	dao := testUserCredentialDao{emailExists: true, emailExistsErr: nil}
	service := NewUserRegistrationLiteService(&dao)
	valid, err := service.Validate(credential.Email)
	require.Nil(t, err)
	require.False(t, valid)
}

func TestGivenDaoHasErrorsWhenUserRegistrationLiteServiceValidateThenReturnError(t *testing.T) {
	dao := testUserCredentialDao{emailExistsErr: errors.New("TEST")}
	service := NewUserRegistrationLiteService(&dao)
	_, err := service.Validate(credential.Email)
	require.NotNil(t, err)
}

func TestGivenValidUserCredentialWhenUserRegistrationLiteServiceRegisterThenNoError(t *testing.T) {
	dao := testUserCredentialDao{createErr: nil}
	service := NewUserRegistrationLiteService(&dao)
	require.Nil(t, service.Register(credential))
}

func TestGivenDaoHasErrorsWhenUserRegistrationLiteServiceRegisterThenReturnError(t *testing.T) {
	dao := testUserCredentialDao{createErr: errors.New("TEST")}
	service := NewUserRegistrationLiteService(&dao)
	require.NotNil(t, service.Register(credential))
}

func startRedis(t *testing.T) *miniredis.Miniredis {
	r, err := miniredis.Run()
	require.Nil(t, err)
	return r
}

func createPool(r *miniredis.Miniredis) *redis.Pool {
	return &redis.Pool{Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", r.Addr())
	}}
}

func TestGivenEmailNotInRedisWhenUserRegistrationLiteServiceValidateThenReturnTrue(t *testing.T) {
	r := startRedis(t)
	defer r.Close()
	pool := createPool(r)
	defer pool.Close()
	dao := dao.NewDefaultUserCredentialDao(pool)
	service := NewUserRegistrationLiteService(dao)
	valid, err := service.Validate(credential.Email)
	require.Nil(t, err)
	require.True(t, valid)
}

func TestGivenEmailInRedisWhenUserRegistrationLiteServiceValidateThenReturnFalse(t *testing.T) {
	r := startRedis(t)
	defer r.Close()
	pool := createPool(r)
	defer pool.Close()
	dao := dao.NewDefaultUserCredentialDao(pool)
	dao.Create(credential)
	service := NewUserRegistrationLiteService(dao)
	valid, err := service.Validate(credential.Email)
	require.Nil(t, err)
	require.False(t, valid)
}

func TestGivenUserCredentialNotInRedisWhenUserRegistrationLiteServiceRegisterThenAddToStore(t *testing.T) {
	r := startRedis(t)
	defer r.Close()
	pool := createPool(r)
	defer pool.Close()
	dao := dao.NewDefaultUserCredentialDao(pool)
	service := NewUserRegistrationLiteService(dao)
	require.Nil(t, service.Register(credential))
	exists, err := dao.EmailExists(credential.Email)
	require.Nil(t, err)
	require.True(t, exists)
}
