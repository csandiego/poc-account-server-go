package dao

import (
        "fmt"
        "github.com/csandiego/poc-account-server/data"
        "github.com/gomodule/redigo/redis"
)

const (
        userCredentialIdCounterKey = "user_credential_id_counter"
        userCredentialKeyFmt = "user_credential:%d"
)


type DefaultUserCredentialDao struct {
        pool *redis.Pool
}

func NewDefaultUserCredentialDao(pool *redis.Pool) *DefaultUserCredentialDao {
        return &DefaultUserCredentialDao{pool}
}

func (dao *DefaultUserCredentialDao) Insert(credential data.UserCredential) error {
        conn := dao.pool.Get()
        defer conn.Close()
        if credential.Id == 0 {
                id, err := redis.Int(conn.Do("INCR", userCredentialIdCounterKey))
                if err != nil {
                        return err
                }
               credential.Id = id
        } else {
                id, err := redis.Int(conn.Do("GET", userCredentialIdCounterKey))
                if err != nil && err != redis.ErrNil {
                        return err
                }
                if credential.Id > id {
                        if _, err = conn.Do("SET", userCredentialIdCounterKey, credential.Id); err != nil {
                                return err
                        }
                }
        }
        args := redis.Args{}.Add(fmt.Sprintf(userCredentialKeyFmt, credential.Id)).AddFlat(credential)
        _, err := conn.Do("HMSET", args...)
        return err
}

func (dao *DefaultUserCredentialDao) Get(id int) (*data.UserCredential, error) {
        conn := dao.pool.Get()
        defer conn.Close()
        reply, err := redis.Values(conn.Do("HGETALL", fmt.Sprintf(userCredentialKeyFmt, id)))
        if err != nil {
                return nil, err
        }
        credential := &data.UserCredential{}
        if err = redis.ScanStruct(reply, credential); err != nil {
                return nil, err
        }
        return credential, nil
}

func (dao *DefaultUserCredentialDao) Update(credential data.UserCredential) error {
        conn := dao.pool.Get()
        defer conn.Close()
        args := redis.Args{}.Add(fmt.Sprintf(userCredentialKeyFmt, credential.Id)).AddFlat(credential)
        _, err := conn.Do("HMSET", args...)
        return err
}

func (dao *DefaultUserCredentialDao) Delete(id int) error {
        conn := dao.pool.Get()
        defer conn.Close()
        _, err := conn.Do("DEL", fmt.Sprintf(userCredentialKeyFmt, id))
        return err
}
