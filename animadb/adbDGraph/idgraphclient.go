package adbDGraph

import (
	"github.com/dgraph-io/dgraph/client"
)

// IDGraphClient - Interface For Interacting With A DGraph Implementation
type IDGraphClient interface {

	// Methods For Connecting To The DGraph Client
	Start(DGraphSettings) error
	Stop() error
	Restart() error
	// Methods For Applying DGraph Settings
	GetSettings() DGraphSettings
	ApplyNewSettings(DGraphSettings) error
	// Methods For Interacting With DGraph Data Processes
	StartTransaction() (*client.Txn, error)
	CommitTransaction(*client.Txn) error
}
