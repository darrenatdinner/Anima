package provision

import (
	"fmt"
	"reflect"
)

// DGraphCommands provides an explicit interface to send terminal commands
type DGraphCommands struct {
	// Flags ----------------------------
	Bulk   bool `cmd:"dgraph bulk"`
	Live   bool `cmd:"dgraph live"`
	Server bool `cmd:"dgraph server"`
	Zero   bool `cmd:"dgraph zero"`
	// Global Flags ---------------------

	// Enable profiling mode, one of [cpu, mem, mutex, block]
	ProfileMode string `cmd:"--profile_mode"`
	// Block profiling rate. Must be used along with block profile_mode
	BlockRate int `cmd:"--block_rate"`
	// Configuration file. Takes precedence over default values, but is overridden to values set with environment variables and flags.
	Config string `cmd:"--config"`
}

// DGraphZeroCommands provides an explicit interface to send terminal commands
type DGraphZeroCommands struct {
	// Flags ----------------------------

	// Use 0.0.0.0 instead of localhost to bind to all addresses on local machine
	Bindall bool `cmd:"--bindall"`
	// Unique node index for this server. (default 1)
	IDX uint `cmd:"--idx"`
	// addr:port of this server, so other Dgraph servers can talk to this
	My string `cmd:"--my"`
	// Address of another dgraphzero server
	Peer string `cmd:"--peer"`
	// Value added to all listening port numbers. [Grpc=7080, HTTP=8080]
	PortOffset int `cmd:"--port_offset"`
	// How many replicas to run per data shard. The count includes the original shard. (default 1)
	Replicas int `cmd:"--replicas"`
	// Directory storing WAL. (default "zw")
	WAL string `cmd:"--wal"`

	// Global Flags ---------------------

	// Enable profiling mode, one of [cpu, mem, mutex, block]
	ProfileMode string `cmd:"--profile_mode"`
	// Block profiling rate. Must be used along with block profile_mode
	BlockRate int `cmd:"--block_rate"`
	// Configuration file. Takes precedence over default values, but is overridden to values set with environment variables and flags.
	Config string `cmd:"--config"`
}

// GenerateCommand creates a string compatible with the DGraph Zero Terminal start command
func GenerateCommand(cmds interface{}) (cmd string) {

	// TODO Use Proper Reflection To Generate Type Dynamically And Get Default Values
	defaultValues := reflect.ValueOf(DGraphZeroCommands{})
	reflectValue := reflect.ValueOf(cmds)
	propertyCount := reflectValue.NumField()
	currentCommand := ""

	for i := 0; i < propertyCount; i++ {
		// If The Command Hasn't Been Specified, Don't Bother
		if reflectValue.Field(i).Interface() == defaultValues.Field(i).Interface() {
			continue
		}
		// Current Value Of The Field
		val := reflectValue.Field(i).Interface()
		// Specified Command Of The Current Field
		command := reflect.TypeOf(cmds).Field(i).Tag.Get("cmd")
		currentCommand = fmt.Sprintf("%s %v", command, val)
		// Add A Space Before Commands
		if cmd != "" {
			cmd += " "
		}
		// Concat The Command
		cmd = cmd + currentCommand
	}

	return
}
