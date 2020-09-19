package main

import (
	pb "github.com/csandiego/poc-account-server/protobuf"
	"google.golang.org/grpc"
	"log"
	"context"
)

func main() {
	conn, err := grpc.Dial("192.168.2.12:8000", grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserRegistrationClient(conn)
	email := "xyz"
	// password := "abc"
	// cred := pb.UserCredential{Email: &email, Password: &password}
	// resp, err := client.Authenticate(context.Background(), &cred)
	req := pb.ValidationRequest{Email: &email}
	_, err = client.Validate(context.Background(), &req)
		if err != nil {
		log.Fatalf("Failed: %v", err)
	}
	_, err = client.Validate(context.Background(), &req)
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}
}
