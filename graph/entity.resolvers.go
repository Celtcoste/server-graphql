package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	generated1 "github.com/Celtcoste/server-graphql/graph/generated"
	model1 "github.com/Celtcoste/server-graphql/graph/model"
)

// Entity returns generated1.EntityResolver implementation.
func (r *Resolver) Entity() generated1.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
