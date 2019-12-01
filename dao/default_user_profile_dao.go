package dao

import (
        "database/sql"
        "github.com/csandiego/poc-account-server/data"
)

type DefaultUserProfileDao struct {
        db *sql.DB
}

func NewDefaultUserProfileDao(db *sql.DB) *DefaultUserProfileDao {
        return &DefaultUserProfileDao{db}
}

func (dao *DefaultUserProfileDao) insert(tx *sql.Tx, profile data.UserProfile) (sql.Result, error) {
        return tx.Exec("INSERT INTO user_profiles (first_name, last_name, date_of_birth) VALUES (?, ?, ?)",
                profile.FirstName, profile.LastName, profile.DateOfBirth)
}

func (dao *DefaultUserProfileDao) get(tx *sql.Tx, id int) (*data.UserProfile, error) {
        profile := &data.UserProfile{}
        if err := tx.QueryRow("SELECT id, first_name, last_name, date_of_birth FROM user_profiles WHERE id = ?", id).
                Scan(&profile.Id, &profile.FirstName, &profile.LastName, &profile.DateOfBirth); err != nil {
                return nil, err
        }
        return profile, nil
}

func (dao *DefaultUserProfileDao) update(tx *sql.Tx, profile data.UserProfile) (sql.Result, error) {
        return tx.Exec("UPDATE user_profiles SET first_name = ?, last_name = ?, date_of_birth = ? WHERE id = ?",
                profile.FirstName, profile.LastName, profile.DateOfBirth, profile.Id)
}

func (dao *DefaultUserProfileDao) delete(tx *sql.Tx, id int) (sql.Result, error) {
        return tx.Exec("DELETE FROM user_profiles WHERE id = ?", id)
}
