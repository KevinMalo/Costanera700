package database

import (
	"context"
	"fmt"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

var ctx = context.Background()

var (
	host = "192.168.100.100"
	port ="9080"
)

func NewClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial( net.JoinHostPort(host,port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

func Commit(p []byte) {

	//COMMIT
	dgraphClient := NewClient()

	mu := &api.Mutation{
		CommitNow: true,
	}

	mu.SetJson = p
	assigned, err := dgraphClient.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(assigned)
}
