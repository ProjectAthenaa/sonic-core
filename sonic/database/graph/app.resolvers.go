package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/graph/generated"
)

func (r *appResolver) ID(ctx context.Context, obj *ent.App) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *appResolver) Settings(ctx context.Context, obj *ent.App) (*ent.Settings, error) {
	panic(fmt.Errorf("not implemented"))
}

// App returns generated.AppResolver implementation.
func (r *Resolver) App() generated.AppResolver { return &appResolver{r} }

type appResolver struct{ *Resolver }
