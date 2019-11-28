package data

type UserCredential struct {
        Id int `redis:"id"`
        Email string `redis:"email"`
        Password string `redis:"password"`
}
