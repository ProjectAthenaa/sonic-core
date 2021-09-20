package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/graph/generated"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/graph/model"
)

func (r *queryResolver) GetUserData(ctx context.Context) (*ent.Metadata, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserCheckouts(ctx context.Context, limit *int) ([]*ent.Checkout, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserStats(ctx context.Context) (*model.Statistics, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserLicense(ctx context.Context) (*ent.License, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Stats(ctx context.Context, obj *ent.User) (*model.Statistics, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
