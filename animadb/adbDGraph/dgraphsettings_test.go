package adbDGraph_test

import (
	"projectanima/AnimaDB/adbDGraph"
	"testing"
)

func TestDGraphSettings_GetURI(t *testing.T) {

	dSettings := adbDGraph.DGraphSettings{
		Path: "localhost",
		Port: "8000",
	}

	if dSettings.GetURI() != "localhost:8000" {
		t.Error("URI Not Properly Formatted")
	}
}

func TestDGraphSettings_NewClientFromSettings(t *testing.T) {

	dSettings := adbDGraph.DGraphSettings{
		Path: "localhost",
		Port: "8000",
		ClusterConnection: false,
		MultiInstanceConnection:false,
	}

	// >> GRPCClient
	regClient, _ := dSettings.NewClientFromSettings()
	if regClient == nil {
		t.Error("Incorrect Client Created")
	}
	// >> GRPCClusterClient

	// >> GRPCMultiClient
	dSettings.MultiInstanceConnection = true
	dSettings.ClusterConnection = true
	regClient, _ = dSettings.NewClientFromSettings()
	if regClient != nil {
		t.Error("Incorrect Client Created")
	}
}
