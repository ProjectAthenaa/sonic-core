package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/product"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/graph/generated"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/graph/model"
)

func (r *accountGroupResolver) ID(ctx context.Context, obj *ent.AccountGroup) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountGroupResolver) Site(ctx context.Context, obj *ent.AccountGroup) (product.Site, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountGroupResolver) Accounts(ctx context.Context, obj *ent.AccountGroup) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAccountGroup(ctx context.Context, newAccountGroup model.AccountGroupInput) (*ent.AccountGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAccountGroup(ctx context.Context, accountGroupID string, updatedAccountGroup model.AccountGroupInput) (*ent.AccountGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllAccountGroups(ctx context.Context) ([]*ent.AccountGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

// AccountGroup returns generated.AccountGroupResolver implementation.
func (r *Resolver) AccountGroup() generated.AccountGroupResolver { return &accountGroupResolver{r} }

type accountGroupResolver struct{ *Resolver }
