package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/graph/generated"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/graph/model"
)

func (r *moduleFieldResolver) Type(ctx context.Context, obj *sonic.ModuleField) (model.FieldType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *moduleFieldResolver) Label(ctx context.Context, obj *sonic.ModuleField) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetModules(ctx context.Context) ([]*model.Module, error) {
	panic(fmt.Errorf("not implemented"))
}

// ModuleField returns generated.ModuleFieldResolver implementation.
func (r *Resolver) ModuleField() generated.ModuleFieldResolver { return &moduleFieldResolver{r} }

type moduleFieldResolver struct{ *Resolver }
