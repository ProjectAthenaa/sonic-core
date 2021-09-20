package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/graph/generated"
)

func (r *checkoutResolver) ID(ctx context.Context, obj *ent.Checkout) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *checkoutResolver) CurrentProductPrice(ctx context.Context, obj *ent.Checkout) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Checkout returns generated.CheckoutResolver implementation.
func (r *Resolver) Checkout() generated.CheckoutResolver { return &checkoutResolver{r} }

type checkoutResolver struct{ *Resolver }
