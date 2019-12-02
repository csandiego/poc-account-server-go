package grpc

import (
	"github.com/csandiego/poc-account-server/data"
	"google.golang.org/grpc"
	"net"
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

func startServer(server *grpc.Server) (net.Listener, *grpc.ClientConn, error) {
	listener, err := net.Listen("tcp", "localhost:")
	if err != nil {
		return nil, nil, err
	}
	go func() {
		server.Serve(listener)
	}()
	conn, err := grpc.Dial(listener.Addr().String(), grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		server.Stop()
		listener.Close()
		return nil, nil, err
	}
	return listener, conn, nil
}

func stopServer(server *grpc.Server, listener net.Listener, conn *grpc.ClientConn) {
	conn.Close()
	server.Stop()
	listener.Close()
}
