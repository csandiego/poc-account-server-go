package gqlgen

import (
	"context"

	"github.com/csandiego/poc-account-server/data"
	"github.com/csandiego/poc-account-server/service"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	userProfileService service.UserProfileService
}

func NewResolver(userProfileService service.UserProfileService) *Resolver {
	return &Resolver{userProfileService: userProfileService}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) UserProfile(ctx context.Context, userID int) (*data.UserProfile, error) {
	return r.userProfileService.Get(userID)
}
