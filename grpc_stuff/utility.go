package rpc_stuff

import (
	"context"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	connMap connectionMap
)

func init() {
	// connections should be established here via env vars
	// using newClientConnections
	newConnMap := make(map[string]*grpc.ClientConn)
	var mut sync.RWMutex
	connMap.connMap = newConnMap
	connMap.mut = &mut

	// address, name
	newClientConn(":50051", "server_1")
}

type connectionMap struct {
	connMap map[string]*grpc.ClientConn
	mut     *sync.RWMutex
}

type RpcTasksInterface interface {
	MakeCall(service, name string) (string, error)
}

type RpcTasks struct {
	conns RpcTasksInterface
}

func NewRpcTasks(conns RpcTasksInterface) *RpcTasks {
	return &RpcTasks{
		conns: conns,
	}
}

func (rt *RpcTasks) MakeCall(service, name string) (string, error) {
	connMap.mut.RLock()
	defer connMap.mut.RUnlock()

	c := NewGreeterClient(connMap.connMap[service])

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &HelloRequest{
		Name: name,
	})
	if err != nil {
		return "", err
	}

	return r.GetMessage(), nil
}

func newClientConn(addr, name string) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v\n", err)
		panic("couldn't connect")
	}

	connMap.mut.Lock()
	connMap.connMap[name] = conn
	connMap.mut.Unlock()
}

func ShutDownAllConnections() {
	connMap.mut.Lock()
	defer connMap.mut.Unlock()
	for _, v := range connMap.connMap {
		v.Close()
	}
}
