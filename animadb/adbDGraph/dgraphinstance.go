package adbDGraph

import (
	"context"

	"github.com/dgraph-io/dgraph/client"
	"github.com/dgraph-io/dgraph/protos/api"
	"google.golang.org/grpc"
)

// DGraphInstance - Encapsulation Of A Single Instance Of A Connection To A DGraph Server
type DGraphInstance struct {
	ActiveConnection   *grpc.ClientConn
	ActiveDGraphClient api.DgraphClient
	ActiveDGraph       *client.Dgraph
	CurrentContext     context.Context
}
