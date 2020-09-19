package mongodb

import (
	"context"
	"github.com/csandiego/poc-account-server/data"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os/exec"
	"testing"
)

const (
	mongodPath = "/Users/toppy/mongodb-macos-x86_64-4.2.2/bin/mongod"
	mongoUri   = "mongodb://localhost"
	mongoDb    = "poc-account-server-test"
)

var mongodArgs = []string{"--dbpath", "/Users/toppy/mongodb-macos-x86_64-4.2.2/data/db"}

func createDatabase(t *testing.T) (*mongo.Database, func()) {
	cmd := exec.Command(mongodPath, mongodArgs...)
	err := cmd.Start()
	require.Nil(t, err)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	require.Nil(t, err)
	err = client.Ping(context.Background(), readpref.Primary())
	require.Nil(t, err)
	db := client.Database(mongoDb)
	return db, func() {
		db.Drop(context.Background())
		cmd.Process.Kill()
	}
}

var profile = data.UserProfile{UserId: 1, FirstName: "abc", LastName: "xyz"}
