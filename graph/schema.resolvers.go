package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"main/graph/generated"
	"main/graph/model"
	rpc_stuff "main/grpc_stuff"
)

// ReturnTime is the resolver for the returnTime field.
func (r *queryResolver) ReturnTime(ctx context.Context) (*model.ReturnTime, error) {
	str3 := "How are you"

	return &model.ReturnTime{
		Server3: &str3,
	}, nil
}

// Server1 is the resolver for the server_1 field.
func (r *returnTimeResolver) Server1(ctx context.Context, obj *model.ReturnTime) (*string, error) {
	str1, err := r.Rpc_Conns.MakeCall(":50051", "server_1")
	if err != nil {
		return nil, err
	}
	return &str1, nil
}

// Server2 is the resolver for the server_2 field.
func (r *returnTimeResolver) Server2(ctx context.Context, obj *model.ReturnTime) (*string, error) {
	str2 := rpc_stuff.ClientConn(":50052", "server_2")
	return &str2, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// ReturnTime returns generated.ReturnTimeResolver implementation.
func (r *Resolver) ReturnTime() generated.ReturnTimeResolver { return &returnTimeResolver{r} }

type queryResolver struct{ *Resolver }
type returnTimeResolver struct{ *Resolver }
