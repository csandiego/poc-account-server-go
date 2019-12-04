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
	server.Validate(context.Background(), &pb.ValidationRequest{Email: &credential.Email})
	require.Equal(t, credential.Email, service.validateEmail)
}

func TestWhenUserRegistrationServerValidateThenWrapResponse(t *testing.T) {
	service := testUserRegistrationService{validateResult: true}
	server := NewUserRegistrationServer(&service)
	response, _ := server.Validate(context.Background(), &pb.ValidationRequest{Email: &credential.Email})
	require.Equal(t, service.validateResult, *response.Valid)
}

func TestWhenUserRegistrationServerRegisterThenCopyParameterFields(t *testing.T) {
	service := testUserRegistrationService{}
	server := NewUserRegistrationServer(&service)
	server.Register(context.Background(), &pb.UserCredential{Email: &credential.Email, Password: &credential.Password})
	require.Equal(t, credential.Email, service.registerCredential.Email)
	require.Equal(t, credential.Password, service.registerCredential.Password)
}
