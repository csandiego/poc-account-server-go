package redigo

import (
	"fmt"
	"github.com/alicebob/miniredis"
	"github.com/csandiego/poc-account-server/data"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

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

var credential = data.UserCredential{Email: "someone@somewhere.com", Password: "password"}

func loadUserCredential(t *testing.T, r *miniredis.Miniredis) {
	id, err := r.Incr(userIdCounterKey, 1)
	require.Nil(t, err)
	key := fmt.Sprintf(userCredentialKeyFmt, credential.Email)
	r.HSet(key, userCredentialPasswordKey, credential.Password)
	r.HSet(key, userCredentialUserIdKey, strconv.FormatInt(int64(id), 10))
}
