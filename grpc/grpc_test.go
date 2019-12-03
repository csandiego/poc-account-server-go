package grpc

import (
	"github.com/csandiego/poc-account-server/data"
)

var credential = data.UserCredential{Email: "someone@somewhere.com", Password: "password"}

type testAuthenticationService struct {
	authenticateCredential data.UserCredential
	authenticateUserId     int
	authenticateErr        error
}

func (service *testAuthenticationService) Authenticate(credential data.UserCredential) (int, error) {
	service.authenticateCredential = credential
	return service.authenticateUserId, service.authenticateErr
}

type testUserRegistrationService struct {
	validateEmail      string
	validateResult     bool
	validateErr        error
	registerCredential data.UserCredential
	registerErr        error
}

func (service *testUserRegistrationService) Validate(email string) (bool, error) {
	service.validateEmail = email
	return service.validateResult, service.validateErr
}

func (service *testUserRegistrationService) Register(credential data.UserCredential) error {
	service.registerCredential = credential
	return service.registerErr
}
