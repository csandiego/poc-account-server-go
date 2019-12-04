package grpc

import (
	"github.com/csandiego/poc-account-server/data"
	pb "github.com/csandiego/poc-account-server/protobuf"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGivenMessageWhenFromMessageThenCopyFieldsToStruct(t *testing.T) {
	src := &pb.UserCredential{Email: &credential.Email, Password: &credential.Password}
	dst := &data.UserCredential{}
	require.Nil(t, fromMessage(src, dst))
	require.Equal(t, *src.Email, dst.Email)
	require.Equal(t, *src.Password, dst.Password)
}
