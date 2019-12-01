package dao

import (
        "fmt"
        "github.com/alicebob/miniredis"
        "github.com/csandiego/poc-account-server/data"
        "github.com/gomodule/redigo/redis"
        "strconv"
        "testing"
)

func startRedis(t *testing.T) *miniredis.Miniredis {
        r, err := miniredis.Run()
        if err != nil {
                t.Fatal(err)
        }
        return r
}

func createPool(r *miniredis.Miniredis) *redis.Pool {
        return &redis.Pool{Dial: func() (redis.Conn, error) {
                return redis.Dial("tcp", r.Addr())
        }}
}

var credential = data.UserCredential{1, "someone@somewhere.com", "password"}

func loadUserCredential(r *miniredis.Miniredis) {
        key := fmt.Sprintf(userCredentialKeyFmt, credential.Id)
        r.HSet(key, "id", strconv.FormatInt(int64(credential.Id), 10))
        r.HSet(key, "email", credential.Email)
        r.HSet(key, "password", credential.Password)
}

func TestValidUserCredentialWithoutIdWhenInsertedThenAddToDatabase(t *testing.T) {
        r := startRedis(t)
        defer r.Close()
        pool := createPool(r)
        defer pool.Close()
        credential := data.UserCredential{Email: "someone@somewhere.com", Password: "password"}
        dao := NewDefaultUserCredentialDao(pool)
        if err := dao.Insert(credential); err != nil {
                t.Fatal(err)
        }
        idString, err := r.Get(userCredentialIdCounterKey)
        if err != nil {
                t.Fatal(err)
        }
        id, err := strconv.ParseInt(idString, 10, 64)
        if err != nil {
                t.Fatal(err)
        }
        if !r.Exists(fmt.Sprintf(userCredentialKeyFmt, id)) {
                t.Fatal("Record not inserted")
        }
}

func TestValidUserCredentialWithIdWhenInsertedThenAddToDatabase(t *testing.T) {
        r := startRedis(t)
        defer r.Close()
        pool := createPool(r)
        defer pool.Close()
        credential := data.UserCredential{10, "someone@somewhere.com", "password"}
        dao := NewDefaultUserCredentialDao(pool)
        if err := dao.Insert(credential); err != nil {
                t.Fatal(err)
        }
        idString, err := r.Get(userCredentialIdCounterKey)
        if err != nil {
                t.Fatal(err)
        }
        id, err := strconv.ParseInt(idString, 10, 64)
        if err != nil {
                t.Fatal(err)
        }
        if id != int64(credential.Id) {
                t.Fatal("ID counter not set")
        }
        if !r.Exists(fmt.Sprintf(userCredentialKeyFmt, id)) {
                t.Fatal("Record not inserted")
        }
}

func TestGivenValidUserIdWhenGetThenFetchUserCredential(t *testing.T) {
        r := startRedis(t)
        defer r.Close()
        pool := createPool(r)
        defer pool.Close()
        loadUserCredential(r)
        dao := NewDefaultUserCredentialDao(pool)
        fetched, err := dao.Get(credential.Id)
        if err != nil {
                t.Fatal(err)
        }
        if credential != *fetched {
                t.Fatal("Record not selected")
        }
}

func TestGivenValidUserCredentialWhenUpdateThenUpdateDatabase(t *testing.T) {
        r := startRedis(t)
        defer r.Close()
        pool := createPool(r)
        defer pool.Close()
        loadUserCredential(r)
        edited := credential
        edited.Email = "edited@somewhere.com"
        dao := NewDefaultUserCredentialDao(pool)
        if err := dao.Update(edited); err != nil {
                t.Fatal(err)
        }
        if r.HGet(fmt.Sprintf(userCredentialKeyFmt, credential.Id), "email") != edited.Email {
                t.Fatal("Record not updated")
        }
}

func TestGivenValidUserIdWhenDeleteThenRemoveFromDatabase(t *testing.T) {
        r := startRedis(t)
        defer r.Close()
        pool := createPool(r)
        defer pool.Close()
        loadUserCredential(r)
        dao := NewDefaultUserCredentialDao(pool)
        if err := dao.Delete(credential.Id); err != nil {
                t.Fatal(err)
        }
        if r.Exists(fmt.Sprintf(userCredentialKeyFmt, credential.Id)) {
                t.Fatal("Record not deleted")
        }
}
