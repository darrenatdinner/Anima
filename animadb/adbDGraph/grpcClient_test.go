package adbDGraph_test

import (
	"projectanima/AnimaDB/adbDGraph"
	"testing"

	"google.golang.org/grpc"
)

func TestGRPCClient_Start(t *testing.T) {

	clientInstance := adbDGraph.GRPCClient{}
	clientSettings := adbDGraph.DGraphSettings{
		ClusterConnection:       false,
		MultiInstanceConnection: false,
		Path:             "localhost",
		Port:             "7080",
		DGraphDialOption: grpc.WithInsecure(),
	}
	errorFound := clientInstance.Start(clientSettings)
	clientInstance.Stop()

	if errorFound != nil {
		t.Error("Start Failed")
	}
}

func TestGRPCClient_Restart(t *testing.T) {

	clientInstance := adbDGraph.GRPCClient{}
	clientSettings := adbDGraph.DGraphSettings{
		ClusterConnection:       false,
		MultiInstanceConnection: false,
		Path:             "localhost",
		Port:             "7080",
		DGraphDialOption: grpc.WithInsecure(),
	}
	_ = clientInstance.Start(clientSettings)

	errorFound := clientInstance.Restart()

	clientInstance.Stop()

	if errorFound != nil {
		t.Error("Restart Failed")
	}
}
func TestGRPCClient_Stop(t *testing.T) {

	clientInstance := adbDGraph.GRPCClient{}
	clientSettings := adbDGraph.DGraphSettings{
		ClusterConnection:       false,
		MultiInstanceConnection: false,
		Path:             "localhost",
		Port:             "7080",
		DGraphDialOption: grpc.WithInsecure(),
	}
	_ = clientInstance.Start(clientSettings)
	errorFound := clientInstance.Stop()

	if errorFound != nil {
		t.Error("Stop Failed")
	}
}
func TestGRPCClient_GetSettings(t *testing.T) {
	clientInstance := adbDGraph.GRPCClient{}
	clientSettings := adbDGraph.DGraphSettings{
		ClusterConnection:       false,
		MultiInstanceConnection: false,
		Path:             "localhost",
		Port:             "7080",
		DGraphDialOption: grpc.WithInsecure(),
	}
	_ = clientInstance.Start(clientSettings)
	clientInstance.Stop()

	if clientSettings.Port != clientInstance.GetSettings().Port {
		t.Error("Get Settings Returned Incorrect Settings")
	}
}
func TestGRPCClient_ApplyNewSettings(t *testing.T) {
	clientInstance := adbDGraph.GRPCClient{}
	clientSettings := adbDGraph.DGraphSettings{
		ClusterConnection:       false,
		MultiInstanceConnection: false,
		Path:             "localhost",
		Port:             "7080",
		DGraphDialOption: grpc.WithInsecure(),
	}
	_ = clientInstance.Start(clientSettings)
	clientSettings2 := adbDGraph.DGraphSettings{
		ClusterConnection:       false,
		MultiInstanceConnection: false,
		Path:             "localhostess",
		Port:             "6080",
		DGraphDialOption: grpc.WithInsecure(),
	}
	_ = clientInstance.ApplyNewSettings(clientSettings2)
	clientInstance.Stop()

	if clientInstance.GetSettings().Port == clientSettings.Port {
		t.Error("Client Setting Not Updated")
	}
}

func TestGRPCClient_StartTransaction(t *testing.T) {

	clientInstance := adbDGraph.GRPCClient{}
	clientSettings := adbDGraph.DGraphSettings{
		ClusterConnection:       false,
		MultiInstanceConnection: false,
		Path:             "localhost",
		Port:             "7080",
		DGraphDialOption: grpc.WithInsecure(),
	}
	_ = clientInstance.Start(clientSettings)

	txn, errorFound := clientInstance.StartTransaction()
	defer txn.Discard(clientInstance.CurrentContext)

	if errorFound != nil {
		t.Error("Start Failed")
	}
	clientInstance.Stop()
}
func TestGRPCClient_CommitTransaction(t *testing.T) {

	clientInstance := adbDGraph.GRPCClient{}
	clientSettings := adbDGraph.DGraphSettings{
		ClusterConnection:       false,
		MultiInstanceConnection: false,
		Path:             "localhost",
		Port:             "7080",
		DGraphDialOption: grpc.WithInsecure(),
	}
	_ = clientInstance.Start(clientSettings)

	txn, _ := clientInstance.StartTransaction()

	errorFound := clientInstance.CommitTransaction(txn)
	if errorFound != nil {
		t.Error("Start Failed")
	}
	clientInstance.Stop()
}

// Negative Scenarios

func TestGRPCClient_ApplyNewSettingsWhileStopped(t *testing.T) {
	clientInstance := adbDGraph.GRPCClient{}
	clientSettings := adbDGraph.DGraphSettings{
		ClusterConnection:       false,
		MultiInstanceConnection: false,
		Path:             "localhost",
		Port:             "7080",
		DGraphDialOption: grpc.WithInsecure(),
	}
	_ = clientInstance.Start(clientSettings)
	clientInstance.Stop()

	clientSettings2 := adbDGraph.DGraphSettings{
		ClusterConnection:       false,
		MultiInstanceConnection: false,
		Path:             "localhostess",
		Port:             "6080",
		DGraphDialOption: grpc.WithInsecure(),
	}
	errorFound := clientInstance.ApplyNewSettings(clientSettings2)
	if errorFound == nil {
		t.Error("No Running Instance To Update Not Thrown")
	}
}
func TestGRPCClient_RestartWhileStopped(t *testing.T) {
	clientInstance := adbDGraph.GRPCClient{}
	clientSettings := adbDGraph.DGraphSettings{
		ClusterConnection:       false,
		MultiInstanceConnection: false,
		Path:             "localhost",
		Port:             "7080",
		DGraphDialOption: grpc.WithInsecure(),
	}
	_ = clientInstance.Start(clientSettings)
	clientInstance.Stop()

	errorFound := clientInstance.Restart()
	if errorFound == nil {
		t.Error("No Running Instance To Update Not Thrown")
	}
}
func TestGRPCClient_StopWhileStopped(t *testing.T) {
	clientInstance := adbDGraph.GRPCClient{}
	errorFound := clientInstance.Stop()

	if errorFound == nil {
		t.Error("No Running Instance To Update Not Thrown")
	}
}

func TestGRPCClient_TransactionWhileStopped(t *testing.T) {
	clientInstance := adbDGraph.GRPCClient{}

	_, errorFound := clientInstance.StartTransaction()

	if errorFound == nil {
		t.Error("No Running Instance To Start Transition Not Thrown")
	}
}
func TestGRPCClient_CommitTransactionWhileStopped(t *testing.T) {
	clientInstance := adbDGraph.GRPCClient{}
	clientSettings := adbDGraph.DGraphSettings{
		ClusterConnection:       false,
		MultiInstanceConnection: false,
		Path:             "localhost",
		Port:             "7080",
		DGraphDialOption: grpc.WithInsecure(),
	}
	_ = clientInstance.Start(clientSettings)
	txn, _ := clientInstance.StartTransaction()
	_ = clientInstance.Stop()
	errorFound := clientInstance.CommitTransaction(txn)

	if errorFound == nil {
		t.Error("No Running Instance To Stop Transition Not Thrown")
	}
}
func TestGRPCClient_CommitTransactionParameterError(t *testing.T) {
	clientInstance := adbDGraph.GRPCClient{}
	clientSettings := adbDGraph.DGraphSettings{
		ClusterConnection:       false,
		MultiInstanceConnection: false,
		Path:             "localhost",
		Port:             "7080",
		DGraphDialOption: grpc.WithInsecure(),
	}
	_ = clientInstance.Start(clientSettings)
	txn, _ := clientInstance.StartTransaction()
	defer txn.Discard(clientInstance.CurrentContext)
	errorFound := clientInstance.CommitTransaction(nil)

	if errorFound == nil {
		t.Error("Valid Parameter Not Given Stop Transition Not Thrown")
	}
}
