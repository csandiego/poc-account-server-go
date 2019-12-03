package grpc

import (
	"context"
	pb "github.com/csandiego/poc-account-server/protobuf"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWhenAuthenticationServerAuthenticateThenCopyParameterFields(t *testing.T) {
	service := testAuthenticationService{}
	server := NewAuthenticationServer(&service)
	server.Authenticate(context.Background(), &pb.UserCredential{Email: credential.Email, Password: credential.Password})
	require.Equal(t, credential.Email, service.authenticateCredential.Email)
	require.Equal(t, credential.Password, service.authenticateCredential.Password)
}

func TestWhenAuthenticationServerAuthenticateThenWrapResponse(t *testing.T) {
	userId := 1
	service := testAuthenticationService{authenticateUserId: userId}
	server := NewAuthenticationServer(&service)
	response, _ := server.Authenticate(context.Background(), &pb.UserCredential{})
	require.Equal(t, int64(userId), response.UserId)
}
