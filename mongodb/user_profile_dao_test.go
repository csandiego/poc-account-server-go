package mongodb

import (
	"context"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestGiveInvalidUserProfileWhereUserProfileDaoCreateThenReturnError(t *testing.T) {
	db, tearDown := createDatabase(t)
	defer tearDown()
	_, err := db.Collection(collectionUserProfiles).InsertOne(
		context.Background(),
		bson.D{
			{"_id", profile.UserId},
			{"firstName", profile.FirstName},
			{"lastName", profile.LastName},
		},
	)
	require.Nil(t, err)
	dao := NewUserProfileDao(db)
	require.NotNil(t, dao.Create(profile))
}

func TestGivenValidUserProfileWhenUserProfileDaoCreateThenInsertIntoDatabase(t *testing.T) {
	db, tearDown := createDatabase(t)
	defer tearDown()
	dao := NewUserProfileDao(db)
	require.Nil(t, dao.Create(profile))
	count, err := db.Collection(collectionUserProfiles).CountDocuments(
		context.Background(),
		bson.D{
			{"_id", profile.UserId},
			{"firstName", profile.FirstName},
			{"lastName", profile.LastName},
		},
	)
	require.Nil(t, err)
	require.Equal(t, int64(1), count)
}

func TestGivenInvalidUserIdWhenUserProfileDaoGetThenReturnError(t *testing.T) {
	db, tearDown := createDatabase(t)
	defer tearDown()
	_, err := db.Collection(collectionUserProfiles).InsertOne(
		context.Background(),
		bson.D{
			{"_id", profile.UserId},
			{"firstName", profile.FirstName},
			{"lastName", profile.LastName},
		},
	)
	require.Nil(t, err)
	dao := NewUserProfileDao(db)
	_, err = dao.Get(0)
	require.NotNil(t, err)
}

func TestGivenValidUserIdWhenUserProfileDaoGetThenFetchFromDatabase(t *testing.T) {
	db, tearDown := createDatabase(t)
	defer tearDown()
	_, err := db.Collection(collectionUserProfiles).InsertOne(
		context.Background(),
		bson.D{
			{"_id", profile.UserId},
			{"firstName", profile.FirstName},
			{"lastName", profile.LastName},
		},
	)
	require.Nil(t, err)
	dao := NewUserProfileDao(db)
	p, err := dao.Get(profile.UserId)
	require.Nil(t, err)
	require.Equal(t, profile, *p)
}

func TestGivenInvalidUserProfileWhenUserProfileDaoUpdateThenReturnError(t *testing.T) {
	db, tearDown := createDatabase(t)
	defer tearDown()
	_, err := db.Collection(collectionUserProfiles).InsertOne(
		context.Background(),
		bson.D{
			{"_id", profile.UserId},
			{"firstName", profile.FirstName},
			{"lastName", profile.LastName},
		},
	)
	require.Nil(t, err)
	edited := profile
	edited.UserId = 0
	dao := NewUserProfileDao(db)
	require.NotNil(t, dao.Update(edited))
}

func TestGivenValidUserProfileWhenUserProfileDaoUpdateThenUpdateDatabase(t *testing.T) {
	db, tearDown := createDatabase(t)
	defer tearDown()
	_, err := db.Collection(collectionUserProfiles).InsertOne(
		context.Background(),
		bson.D{
			{"_id", profile.UserId},
			{"firstName", profile.FirstName},
			{"lastName", profile.LastName},
		},
	)
	require.Nil(t, err)
	edited := profile
	edited.FirstName = "123"
	dao := NewUserProfileDao(db)
	require.Nil(t, dao.Update(edited))
	count, err := db.Collection(collectionUserProfiles).CountDocuments(
		context.Background(),
		bson.D{
			{"_id", profile.UserId},
			{"firstName", profile.FirstName},
			{"lastName", profile.LastName},
		},
	)
	require.Nil(t, err)
	require.Equal(t, int64(0), count)
	count, err = db.Collection(collectionUserProfiles).CountDocuments(
		context.Background(),
		bson.D{
			{"_id", edited.UserId},
			{"firstName", edited.FirstName},
			{"lastName", edited.LastName},
		},
	)
	require.Nil(t, err)
	require.Equal(t, int64(1), count)
}
