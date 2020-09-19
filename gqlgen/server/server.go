package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/csandiego/poc-account-server/data"
	"github.com/csandiego/poc-account-server/gqlgen"
	"github.com/csandiego/poc-account-server/service"
	"net/http"
)

type userProfileDao struct{}

func (dao *userProfileDao) Create(profile data.UserProfile) error {
	return nil
}

func (dao *userProfileDao) Get(userId int) (*data.UserProfile, error) {
	return &data.UserProfile{UserId: userId, FirstName: "abc", LastName: "xyz"}, nil
}

func (dao *userProfileDao) Update(profile data.UserProfile) error {
	return nil
}

func main() {
	service := service.NewDefaultUserProfileService(&userProfileDao{})
	resolver := gqlgen.NewResolver(service)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
        http.Handle("/query", handler.GraphQL(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: resolver})))
	http.ListenAndServe(":8000", nil)
}
