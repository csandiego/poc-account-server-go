package mongodb

import (
	"context"
	"errors"
	"github.com/csandiego/poc-account-server/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionUserProfiles = "user_profiles"

type UserProfileDao struct {
	db *mongo.Database
}

func NewUserProfileDao(db *mongo.Database) *UserProfileDao {
	return &UserProfileDao{db: db}
}

func (dao *UserProfileDao) Create(profile data.UserProfile) error {
	document, err := bson.Marshal(&profile)
	if err != nil {
		return err
	}
	_, err = dao.db.Collection(collectionUserProfiles).InsertOne(
		context.Background(), document,
	)
	return err
}

func (dao *UserProfileDao) Get(userId int) (*data.UserProfile, error) {
	profile := &data.UserProfile{}
	err := dao.db.Collection(collectionUserProfiles).FindOne(
		context.Background(), bson.D{{"_id", userId}},
	).Decode(profile)
	return profile, err
}

func (dao *UserProfileDao) Update(profile data.UserProfile) error {
	document, err := bson.Marshal(&profile)
	if err != nil {
		return err
	}
	result, err := dao.db.Collection(collectionUserProfiles).ReplaceOne(
		context.Background(), bson.D{{"_id", profile.UserId}}, document,
	)
	if err == nil && (result.MatchedCount != int64(1) || result.ModifiedCount != int64(1)) {
		return errors.New("user profile not updated")
	}
	return err
}
