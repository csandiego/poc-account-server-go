package grpc

import (
	"context"
	pb "github.com/csandiego/poc-account-server/protobuf"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWhenUserRegistrationServerValidateThenExtractEmail(t *testing.T) {
	service := testUserRegistrationService{}
	server := NewUserRegistrationServer(&service)
	server.Validate(context.Background(), &pb.ValidationRequest{Email: credential.Email})
	require.Equal(t, credential.Email, service.validateEmail)
}

func TestWhenUserRegistrationServerAuthenticateThenWrapResponse(t *testing.T) {
	valid := true
	service := testUserRegistrationService{validateResult: valid}
	server := NewUserRegistrationServer(&service)
	response, _ := server.Validate(context.Background(), &pb.ValidationRequest{})
	require.Equal(t, valid, response.Valid)
}

func TestWhenUserRegistrationServerRegisterThenCopyParameterFields(t *testing.T) {
	service := testUserRegistrationService{}
	server := NewUserRegistrationServer(&service)
	server.Register(context.Background(), &pb.UserCredential{Email: credential.Email, Password: credential.Password})
	require.Equal(t, credential.Email, service.registerCredential.Email)
	require.Equal(t, credential.Password, service.registerCredential.Password)
}
