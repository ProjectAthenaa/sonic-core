package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/graph/generated"
)

func (r *statisticResolver) ID(ctx context.Context, obj *ent.Statistic) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *statisticResolver) Axis(ctx context.Context, obj *ent.Statistic) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

// Statistic returns generated.StatisticResolver implementation.
func (r *Resolver) Statistic() generated.StatisticResolver { return &statisticResolver{r} }

type statisticResolver struct{ *Resolver }
