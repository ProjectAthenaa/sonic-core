package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/graph/generated"
)

func (r *licenseResolver) ID(ctx context.Context, obj *ent.License) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stripeResolver) ID(ctx context.Context, obj *ent.Stripe) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// License returns generated.LicenseResolver implementation.
func (r *Resolver) License() generated.LicenseResolver { return &licenseResolver{r} }

// Stripe returns generated.StripeResolver implementation.
func (r *Resolver) Stripe() generated.StripeResolver { return &stripeResolver{r} }

type licenseResolver struct{ *Resolver }
type stripeResolver struct{ *Resolver }
