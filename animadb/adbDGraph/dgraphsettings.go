package adbDGraph

import (
	"fmt"

	"google.golang.org/grpc"
)

// DGraphSettings - Settings Object To Encapsulate Various Options To Load DGraph With
type DGraphSettings struct {
	Path                    string
	Port                    string
	ClusterConnection       bool
	MultiInstanceConnection bool
	DGraphDialOption        grpc.DialOption
}

// GetURI - Get The Combined URI Of
func (dgsObject DGraphSettings) GetURI() string {
	return dgsObject.Path + ":" + dgsObject.Port
}

// NewClientFromSettings - Create A New Instance Of The DGraph Client Based On The Interpreted Settings
func (dgsObject DGraphSettings) NewClientFromSettings() (IDGraphClient, error) {

	// ClusterClient: False + MultiInstanceConnection: False => GRPCClient
	if !dgsObject.ClusterConnection && !dgsObject.MultiInstanceConnection {
		newClient := new(GRPCClient)
		newClient.currentSettings = dgsObject
		return newClient, nil
	}
	// ClusterClient: True + MultiInstanceConnection: False => GRPCClusterClient

	// ClusterClient: False + MultiInstanceConnection: True => GRPCMultiClient

	// ClusterClient: True + MultiInstanceConnection: True => ?
	return nil, fmt.Errorf("no connection can be generated with the current settings")
}
