package graph

import rpc_stuff "main/grpc_stuff"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Rpc_Conns rpc_stuff.RpcTasksInterface
}
