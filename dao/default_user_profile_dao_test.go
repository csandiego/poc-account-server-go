package dao

import (
        "database/sql"
        "github.com/csandiego/user-account-server/data"
        _ "github.com/mattn/go-sqlite3"
        "io/ioutil"
        "testing"
        "time"
)

const schemaFile = "../schema.sql"

func createSchema(t *testing.T, db *sql.DB) {
        schema, err := ioutil.ReadFile(schemaFile)
        if err != nil {
                t.Fatal(err)
        }
        if _, err := db.Exec(string(schema)); err != nil {
                t.Fatal(err)
        }
}

var profile = data.UserProfile{1, "First", "Last", time.Now().Truncate(time.Hour * 24)}

func loadData(t *testing.T, db *sql.DB) {
        if _, err := db.Exec("INSERT INTO user_profiles (id, first_name, last_name, date_of_birth) VALUES (?, ?, ?, ?)",
                profile.Id, profile.FirstName, profile.LastName, profile.DateOfBirth); err != nil {
                t.Fatal(err)
        }
}


func createDatabase(t *testing.T) *sql.DB {
        db, err := sql.Open("sqlite3", "file:_?mode=memory&_loc=auto")
        if err != nil {
                t.Fatal(err)
        }
        createSchema(t, db)
        loadData(t, db)
        return db
}

func TestGivenValidUserProfileWhenInsertedThenAddToDatabase(t *testing.T) {
        db := createDatabase(t)
        defer db.Close()
        dao := NewDefaultUserProfileDao(db)
        tx, err := db.Begin()
        if err != nil {
                t.Fatal(err)
        }
        profile := data.UserProfile{FirstName: "First", LastName: "Last", DateOfBirth: time.Now()}
        result, err := dao.insert(tx, &profile)
        if err != nil {
                t.Error(err)
        }
        rows, err := result.RowsAffected()
        if err != nil {
                t.Error(err)
        }
        if rows != 1 {
                t.Error("Record not inserted")
        }
}

func TestGivenValidUserIdWhenGetThenFetchUserProfile(t *testing.T) {
        db := createDatabase(t)
        defer db.Close()
        dao := NewDefaultUserProfileDao(db)
        tx, err := db.Begin()
        if err != nil {
                t.Fatal(err)
        }
        fetched, err := dao.get(tx, profile.Id)
        if err != nil {
                t.Error(err)
        }
        if profile != *fetched {
                t.Error("Record not selected")
        }
}

func TestGivenValidUserProfileWhenUpdateThenUpdateDatabase(t *testing.T) {
        db := createDatabase(t)
        defer db.Close()
        editedProfile := profile
        editedProfile.FirstName = "Edited"
        dao := NewDefaultUserProfileDao(db)
        tx, err := db.Begin()
        if err != nil {
                t.Fatal(err)
        }
        result, err := dao.update(tx, &editedProfile)
        if err != nil {
                t.Error(err)
        }
        rows, err := result.RowsAffected()
        if err != nil {
                t.Error(err)
        }
        if rows != 1 {
                t.Error("Record not updated")
        }
}

func TestGivenValidUserProfileIdWhenDeleteThenRemoveFromDatabase(t *testing.T) {
        db := createDatabase(t)
        defer db.Close()
        dao := NewDefaultUserProfileDao(db)
        tx, err := db.Begin()
        if err != nil {
                t.Fatal(err)
        }
        result, err := dao.delete(tx, profile.Id)
        if err != nil {
                t.Error(err)
        }
        rows, err := result.RowsAffected()
        if err != nil {
                t.Error(err)
        }
        if rows != 1 {
                t.Error("Record not deleted")
        }
}
