package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/graph/generated"
)

func (r *metadataResolver) ID(ctx context.Context, obj *ent.Metadata) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Metadata returns generated.MetadataResolver implementation.
func (r *Resolver) Metadata() generated.MetadataResolver { return &metadataResolver{r} }

type metadataResolver struct{ *Resolver }
