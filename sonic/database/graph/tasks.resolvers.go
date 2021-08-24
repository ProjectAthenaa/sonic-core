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

func (r *mutationResolver) CreateTask(ctx context.Context, newTask model.NewTask) (*ent.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTask(ctx context.Context, taskID string, updatedTask model.UpdatedTask) (*ent.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTask(ctx context.Context, taskID string, deletedProduct bool) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTaskGroup(ctx context.Context, newTaskGroup model.NewTaskGroup) (*ent.TaskGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTaskGroup(ctx context.Context, taskGroupID string, updatedTaskGroup model.NewTaskGroup) (*ent.TaskGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTaskGroup(ctx context.Context, taskGroupID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, productID string, updatedProduct model.ProductIn) (*ent.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) ID(ctx context.Context, obj *ent.Product) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) Metadata(ctx context.Context, obj *ent.Product) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTask(ctx context.Context, taskID string) (*ent.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTaskGroup(ctx context.Context, taskGroupID string) (*ent.TaskGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProduct(ctx context.Context, productID string) (*ent.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllTaskGroups(ctx context.Context) ([]*ent.TaskGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllTasks(ctx context.Context, taskGroupID string) ([]*ent.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *taskResolver) ID(ctx context.Context, obj *ent.Task) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *taskResolver) Product(ctx context.Context, obj *ent.Task) (*ent.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *taskResolver) ProxyList(ctx context.Context, obj *ent.Task) (*ent.ProxyList, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *taskResolver) ProfileGroup(ctx context.Context, obj *ent.Task) (*ent.ProfileGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *taskGroupResolver) ID(ctx context.Context, obj *ent.TaskGroup) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

// TaskGroup returns generated.TaskGroupResolver implementation.
func (r *Resolver) TaskGroup() generated.TaskGroupResolver { return &taskGroupResolver{r} }

type productResolver struct{ *Resolver }
type taskResolver struct{ *Resolver }
type taskGroupResolver struct{ *Resolver }
