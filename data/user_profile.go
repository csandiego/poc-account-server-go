package data

import "time"

type UserProfile struct {
        Id int
        FirstName string
        LastName string
        DateOfBirth time.Time
}
