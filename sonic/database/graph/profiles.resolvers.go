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

func (r *addressResolver) ID(ctx context.Context, obj *ent.Address) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *billingResolver) ID(ctx context.Context, obj *ent.Billing) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProfile(ctx context.Context, newProfile model.NewProfile) (*ent.Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, profileID string, updatedProfile model.NewProfile) (*ent.Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProfile(ctx context.Context, profileID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProfileGroup(ctx context.Context, newGroup model.NewProfileGroup) (*ent.ProfileGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProfileGroup(ctx context.Context, groupID string, updatedGroup model.NewProfileGroup) (*ent.ProfileGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProfileGroup(ctx context.Context, groupID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *profileResolver) ID(ctx context.Context, obj *ent.Profile) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *profileResolver) Shipping(ctx context.Context, obj *ent.Profile) (*ent.Shipping, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *profileResolver) Billing(ctx context.Context, obj *ent.Profile) (*ent.Billing, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *profileGroupResolver) ID(ctx context.Context, obj *ent.ProfileGroup) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProfile(ctx context.Context, profileID string) (*ent.Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProfileGroup(ctx context.Context, profileGroupID string) (*ent.ProfileGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProfileGroups(ctx context.Context) ([]*ent.ProfileGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *shippingResolver) ID(ctx context.Context, obj *ent.Shipping) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Address returns generated.AddressResolver implementation.
func (r *Resolver) Address() generated.AddressResolver { return &addressResolver{r} }

// Billing returns generated.BillingResolver implementation.
func (r *Resolver) Billing() generated.BillingResolver { return &billingResolver{r} }

// Profile returns generated.ProfileResolver implementation.
func (r *Resolver) Profile() generated.ProfileResolver { return &profileResolver{r} }

// ProfileGroup returns generated.ProfileGroupResolver implementation.
func (r *Resolver) ProfileGroup() generated.ProfileGroupResolver { return &profileGroupResolver{r} }

// Shipping returns generated.ShippingResolver implementation.
func (r *Resolver) Shipping() generated.ShippingResolver { return &shippingResolver{r} }

type addressResolver struct{ *Resolver }
type billingResolver struct{ *Resolver }
type profileResolver struct{ *Resolver }
type profileGroupResolver struct{ *Resolver }
type shippingResolver struct{ *Resolver }
