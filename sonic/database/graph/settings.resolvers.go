package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/graph/generated"
)

func (r *mutationResolver) SetSuccessWebhook(ctx context.Context, webhook string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SetDeclineWebhook(ctx context.Context, webhook string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SetCheckoutDelay(ctx context.Context, delay int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SetATCDelay(ctx context.Context, delay int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetSettings(ctx context.Context) (*ent.Settings, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TestSuccessWebhook(ctx context.Context) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TestDeclineWebhook(ctx context.Context) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *settingsResolver) ID(ctx context.Context, obj *ent.Settings) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Settings returns generated.SettingsResolver implementation.
func (r *Resolver) Settings() generated.SettingsResolver { return &settingsResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type settingsResolver struct{ *Resolver }
