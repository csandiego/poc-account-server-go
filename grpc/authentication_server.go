package grpc

import (
	"context"
	"github.com/csandiego/poc-account-server/data"
	pb "github.com/csandiego/poc-account-server/protobuf"
	"github.com/csandiego/poc-account-server/service"
)

type AuthenticationServer struct {
	pb.UnimplementedAuthenticationServer
	service service.AuthenticationService
}

func NewAuthenticationServer(service service.AuthenticationService) *AuthenticationServer {
	return &AuthenticationServer{service: service}
}

func (server *AuthenticationServer) Authenticate(ctx context.Context, req *pb.UserCredential) (*pb.AuthenticationResponse, error) {
	credential := data.UserCredential{}
	fromMessage(req, &credential)
	userId, err := server.service.Authenticate(credential)
	return &pb.AuthenticationResponse{UserId: int64(userId)}, err
}
