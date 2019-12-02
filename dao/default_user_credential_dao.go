package dao

import (
	"errors"
	"fmt"
	"github.com/csandiego/poc-account-server/data"
	"github.com/gomodule/redigo/redis"
)

const (
	userIdCounterKey          = "user_id_counter"
	userCredentialKeyFmt      = "user_credential:%s"
	userCredentialPasswordKey = "password"
	userCredentialUserIdKey   = "user_id"
)

var (
	ErrPasswordMismatch = errors.New("Passwords do not match")
)

type DefaultUserCredentialDao struct {
	pool *redis.Pool
}

func NewDefaultUserCredentialDao(pool *redis.Pool) *DefaultUserCredentialDao {
	return &DefaultUserCredentialDao{pool}
}

func (dao *DefaultUserCredentialDao) EmailExists(email string) (bool, error) {
	conn := dao.pool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("EXISTS", fmt.Sprintf(userCredentialKeyFmt, email)))
}

func (dao *DefaultUserCredentialDao) Create(credential data.UserCredential) error {
	conn := dao.pool.Get()
	defer conn.Close()
	id, err := redis.Int(conn.Do("INCR", userIdCounterKey))
	if err != nil {
		return err
	}
	key := fmt.Sprintf(userCredentialKeyFmt, credential.Email)
	_, err = conn.Do("HMSET", key, userCredentialPasswordKey, credential.Password, userCredentialUserIdKey, id)
	return err
}

type internalUserCredential struct {
	Password string `redis:"password"`
	UserId   int    `redis:"user_id"`
}

func (dao *DefaultUserCredentialDao) Authenticate(credential data.UserCredential) (int, error) {
	conn := dao.pool.Get()
	defer conn.Close()
	reply, err := redis.Values(conn.Do("HGETALL", fmt.Sprintf(userCredentialKeyFmt, credential.Email)))
	if err != nil {
		return 0, err
	}
	holder := internalUserCredential{}
	if err = redis.ScanStruct(reply, &holder); err != nil {
		return 0, err
	}
	if credential.Password != holder.Password {
		return 0, ErrPasswordMismatch
	}
	return holder.UserId, nil
}
