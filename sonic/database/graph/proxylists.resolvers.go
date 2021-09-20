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

func (r *mutationResolver) CreateProxyList(ctx context.Context, proxyList model.NewProxyList) (*ent.ProxyList, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProxyList(ctx context.Context, proxyListID string, proxyList model.NewProxyList) (*ent.ProxyList, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProxyList(ctx context.Context, proxyListID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *proxyResolver) ID(ctx context.Context, obj *ent.Proxy) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *proxyListResolver) ID(ctx context.Context, obj *ent.ProxyList) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProxyList(ctx context.Context, proxyListID string) (*ent.ProxyList, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TestProxyList(ctx context.Context, proxyListID string) ([]*model.ProxyTest, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllProxyLists(ctx context.Context) ([]*ent.ProxyList, error) {
	panic(fmt.Errorf("not implemented"))
}

// Proxy returns generated.ProxyResolver implementation.
func (r *Resolver) Proxy() generated.ProxyResolver { return &proxyResolver{r} }

// ProxyList returns generated.ProxyListResolver implementation.
func (r *Resolver) ProxyList() generated.ProxyListResolver { return &proxyListResolver{r} }

type proxyResolver struct{ *Resolver }
type proxyListResolver struct{ *Resolver }
