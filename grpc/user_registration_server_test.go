package grpc

import (
	"context"
	pb "github.com/csandiego/poc-account-server/protobuf"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"testing"
)

func TestWhenUserRegistrationServerValidateThenExtractEmail(t *testing.T) {
	service := testUserRegistrationService{}
	server := NewUserRegistrationServer(&service)
	grpcServer := grpc.NewServer()
	pb.RegisterUserRegistrationServer(grpcServer, server)
	l, c, err := startServer(grpcServer)
	require.Nil(t, err)
	defer stopServer(grpcServer, l, c)
	client := pb.NewUserRegistrationClient(c)
	_, err = client.Validate(context.Background(), &pb.ValidationRequest{Email: credential.Email})
	require.Nil(t, err)
	require.Equal(t, credential.Email, service.validateEmail)
}

func TestWhenUserRegistrationServerAuthenticateThenWrapResponse(t *testing.T) {
	valid := true
	service := testUserRegistrationService{validateResult: valid}
	server := NewUserRegistrationServer(&service)
	grpcServer := grpc.NewServer()
	pb.RegisterUserRegistrationServer(grpcServer, server)
	l, c, err := startServer(grpcServer)
	require.Nil(t, err)
	defer stopServer(grpcServer, l, c)
	client := pb.NewUserRegistrationClient(c)
	response, err := client.Validate(context.Background(), &pb.ValidationRequest{})
	require.Nil(t, err)
	require.Equal(t, valid, response.Valid)
}

func TestWhenUserRegistrationServerRegisterThenCopyParameterFields(t *testing.T) {
	service := testUserRegistrationService{}
	server := NewUserRegistrationServer(&service)
	grpcServer := grpc.NewServer()
	pb.RegisterUserRegistrationServer(grpcServer, server)
	l, c, err := startServer(grpcServer)
	require.Nil(t, err)
	defer stopServer(grpcServer, l, c)
	client := pb.NewUserRegistrationClient(c)
	_, err = client.Register(context.Background(), &pb.UserCredential{Email: credential.Email, Password: credential.Password})
	require.Nil(t, err)
	require.Equal(t, credential.Email, service.registerCredential.Email)
	require.Equal(t, credential.Password, service.registerCredential.Password)
}
