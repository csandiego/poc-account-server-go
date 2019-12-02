package grpc

import (
	"context"
	"github.com/csandiego/poc-account-server/data"
	pb "github.com/csandiego/poc-account-server/protobuf"
	"github.com/csandiego/poc-account-server/service"
	"github.com/golang/protobuf/ptypes/empty"
)

type UserRegistrationServer struct {
	pb.UnimplementedUserRegistrationServer
	service service.UserRegistrationService
}

func NewUserRegistrationServer(service service.UserRegistrationService) *UserRegistrationServer {
	return &UserRegistrationServer{service: service}
}

func (server *UserRegistrationServer) Validate(ctx context.Context, req *pb.ValidationRequest) (*pb.ValidationResponse, error) {
	valid, err := server.service.Validate(req.Email)
	response := pb.ValidationResponse{Valid: valid}
	return &response, err
}

func (server *UserRegistrationServer) Register(ctx context.Context, req *pb.UserCredential) (*empty.Empty, error) {
	credential := data.UserCredential{}
	fromMessage(req, &credential)
	return &empty.Empty{}, server.service.Register(credential)
}
