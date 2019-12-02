package grpc

import (
	"context"
	pb "github.com/csandiego/poc-account-server/protobuf"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"testing"
)

func TestWhenAuthenticationServerAuthenticateThenCopyParameterFields(t *testing.T) {
	service := testAuthenticationService{}
	server := NewAuthenticationServer(&service)
	grpcServer := grpc.NewServer()
	pb.RegisterAuthenticationServer(grpcServer, server)
	l, c, err := startServer(grpcServer)
	require.Nil(t, err)
	defer stopServer(grpcServer, l, c)
	client := pb.NewAuthenticationClient(c)
	_, err = client.Authenticate(context.Background(), &pb.UserCredential{Email: credential.Email, Password: credential.Password})
	require.Nil(t, err)
	require.Equal(t, credential.Email, service.authenticateCredential.Email)
	require.Equal(t, credential.Password, service.authenticateCredential.Password)
}

func TestWhenAuthenticationServerAuthenticateThenWrapResponse(t *testing.T) {
	userId := 1
	service := testAuthenticationService{authenticateUserId: userId}
	server := NewAuthenticationServer(&service)
	grpcServer := grpc.NewServer()
	pb.RegisterAuthenticationServer(grpcServer, server)
	l, c, err := startServer(grpcServer)
	require.Nil(t, err)
	defer stopServer(grpcServer, l, c)
	client := pb.NewAuthenticationClient(c)
	response, err := client.Authenticate(context.Background(), &pb.UserCredential{})
	require.Nil(t, err)
	require.Equal(t, int64(userId), response.UserId)
}
