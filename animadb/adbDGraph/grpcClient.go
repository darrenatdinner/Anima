package adbDGraph

import (
	"context"
	"fmt"
	"log"

	"github.com/dgraph-io/dgraph/client"
	"github.com/dgraph-io/dgraph/protos/api"
	"google.golang.org/grpc"
)

// GRPCClient - gRPC based client For Managing The Connection With DGraph
//	Instead Of Adding Additional Complexity
type GRPCClient struct {
	// DGraph Implimeted Items
	DGraphInstance
	// Settings Related To The Setup Of A DGraph Instance
	currentSettings DGraphSettings
}

// Start - Method For Connecting To The DGraph Client
func (clientObject *GRPCClient) Start(settings DGraphSettings) error {

	clientObject.currentSettings = settings
	conn, err := grpc.Dial(clientObject.currentSettings.GetURI(), clientObject.currentSettings.DGraphDialOption)
	if err != nil {
		clientObject.ActiveConnection = nil
		log.Fatal("error while trying to dial gRPC")
		return err
	}
	// If Connected Continue With The Process
	clientObject.ActiveConnection = conn
	clientObject.ActiveDGraphClient = api.NewDgraphClient(clientObject.ActiveConnection)
	clientObject.ActiveDGraph = client.NewDgraphClient(clientObject.ActiveDGraphClient)
	// If No Error, Return Nil Notifying Of Correct Setup
	return nil
}

// Stop - Method For Stopping Communication With The Server That Hosts The DGraph Instance
func (clientObject *GRPCClient) Stop() error {
	if clientObject.ActiveConnection == nil {
		return fmt.Errorf("no open connection found")
	}
	clientObject.CurrentContext = nil
	clientObject.ActiveDGraph = nil
	clientObject.ActiveDGraphClient = nil
	return clientObject.ActiveConnection.Close()
}

// Restart - Stop Then           Restart The Client Using The Saved Settings
func (clientObject *GRPCClient) Restart() error {
	errorFound := clientObject.Stop()
	if errorFound != nil {
		return errorFound
	}
	return clientObject.Start(clientObject.currentSettings)
}

// GetSettings - Methods For Applying DGraph Settings
func (clientObject GRPCClient) GetSettings() DGraphSettings {
	return clientObject.currentSettings
}

// ApplyNewSettings - Using A Setting Object, Apply Then Restart The Connection To The Instance
func (clientObject *GRPCClient) ApplyNewSettings(newSettings DGraphSettings) error {
	// Stop
	errorFound := clientObject.Stop()
	if errorFound != nil {
		return errorFound
	}
	// Update
	return clientObject.Start(newSettings)

}

// StartTransaction - Create A Transaction Object To Begin An ACID Transaction
func (clientObject *GRPCClient) StartTransaction() (*client.Txn, error) {
	if clientObject.ActiveDGraph == nil {
		return nil, fmt.Errorf("no active connection to the DGraph instance")
	}
	retTxn := clientObject.ActiveDGraph.NewTxn()
	clientObject.CurrentContext = context.Background()
	return retTxn, nil
}

// CommitTransaction - Commit The Transaction To The Database Signifying There Were No Errors,
//	Call "defer txn.Discard" After Starting A Transaction To Verify The Closure Of The Transaction
//	In The Case Of Success A Transaction Will Be Committed And The Defer Will Be Ignored(Functionally)
func (clientObject *GRPCClient) CommitTransaction(transactionInstance *client.Txn) error {
	if clientObject.ActiveDGraph == nil {
		return fmt.Errorf("no active connection to the DGraph instance")
	}
	if transactionInstance == nil {
		return fmt.Errorf("transaction invalid")
	}
	return transactionInstance.Commit(clientObject.CurrentContext)
}
