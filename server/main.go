package main

import (
	"context"
	"github.com/99designs/gqlgen/handler"
	"github.com/alicebob/miniredis"
	"github.com/csandiego/poc-account-server/gqlgen"
	server "github.com/csandiego/poc-account-server/grpc"
	"github.com/csandiego/poc-account-server/mongodb"
	pb "github.com/csandiego/poc-account-server/protobuf"
	"github.com/csandiego/poc-account-server/redigo"
	"github.com/csandiego/poc-account-server/service"
	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	_ "google.golang.org/grpc/encoding/gzip"
	"log"
	"net"
	"net/http"
)

type plainTextHasher struct {
}

func (h *plainTextHasher) Hash(password string) string {
	return password
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		log.Fatalf("MongoDB: %v\n", err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("MongoDB: %v\n", err)
	}
	db := client.Database("poc-account-server")
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("Miniredis: %v\n", err)
	}
	defer mr.Close()
	pool := &redis.Pool{Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", mr.Addr())
	}}
	defer pool.Close()

	userCredentialDao := redigo.NewRedigoUserCredentialDao(pool)
	userProfileDao := mongodb.NewUserProfileDao(db)
	hasher := &plainTextHasher{}
	userRegistrationService := service.NewDefaultUserRegistrationService(userCredentialDao, hasher)
	authenticationService := service.NewDefaultAuthenticationService(userCredentialDao, hasher)
	userProfileService := service.NewDefaultUserProfileService(userProfileDao)

	go func() {
		userRegistrationServer := server.NewUserRegistrationServer(userRegistrationService)
		authenticationServer := server.NewAuthenticationServer(authenticationService)
		grpcServer := grpc.NewServer()
		pb.RegisterUserRegistrationServer(grpcServer, userRegistrationServer)
		pb.RegisterAuthenticationServer(grpcServer, authenticationServer)
		s, err := net.Listen("tcp", ":8000")
		if err != nil {
			log.Fatalf("net.Listen: %v\n", err)
		}
		defer s.Close()
		encoding.RegisterCompressor(encoding.GetCompressor("gzip"))
		log.Println("Serving GRPC...")
		grpcServer.Serve(s)
	}()

	go func() {
		resolver := gqlgen.NewResolver(userProfileService)
		http.Handle("/", handler.Playground("GraphQL playground", "/query"))
		http.Handle("/graphql", handler.GraphQL(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: resolver})))
		log.Println("Serving GraphQL...")
		http.ListenAndServe(":8080", nil)

	}()

	c := make(chan struct{})
	<-c
}
