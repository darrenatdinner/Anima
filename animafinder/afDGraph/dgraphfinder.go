package afDGraph

import (
	"fmt"
	"projectanima/AnimaDB/adbDGraph"

	"github.com/dgraph-io/dgraph/client"

	"github.com/dgraph-io/dgraph/protos/api"
)

// DGraphFinder - Finder Structure That Works With The DGraph Client To Manipulate,
//	Create, And Delete Data From The Database
type DGraphFinder struct {
	currentClient *adbDGraph.GRPCClient
}

// Query - Send A Query To the Database To Get Data
func (dgfInst DGraphFinder) Query(currTrans *client.Txn, queryString string) (*api.Response, error) {

	if dgfInst.currentClient == nil {
		return nil, fmt.Errorf("internal DGraphClient is invalid")
	}
	queryData := new(api.Request)
	queryData.Query = queryString

	return currTrans.Query(
		dgfInst.currentClient.CurrentContext, queryString)
}

// Mutate - Called Mutate On The DB To Alter Data
func (dgfInst DGraphFinder) Mutate(currTrans *client.Txn, data interface{}) error {

	//return currTrans.
	return nil
}
