package data

type UserProfile struct {
	UserId    int    `bson:"_id"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
}
