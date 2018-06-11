package provision_test

import (
	"projectanima/AnimaDB/adbDGraph/provision"
	"testing"
)

func TestDGraphZeroCommands_GenerateCommand_Default(t *testing.T) {
	brr := provision.DGraphZeroCommands{}

	command := provision.GenerateCommand(brr)

	if command != "" {
		t.Errorf("incorrect command. got: %s, wanted: %s", command, "")
	}
}
func TestDGraphZeroCommands_GenerateCommand_SingleProperty(t *testing.T) {
	brr := provision.DGraphZeroCommands{
		Bindall: true,
	}

	command := provision.GenerateCommand(brr)
	expectedResult := "--bindall true"

	if command == "" {
		t.Errorf("incorrect command. got: %s, wanted: %s", command, "")
	}
	if command != expectedResult {
		t.Errorf("incorrect command. got: %s, wanted: %s", command, expectedResult)
	}
}
func TestDGraphZeroCommands_GenerateCommand_MultiProperty(t *testing.T) {
	brr := provision.DGraphZeroCommands{
		Bindall: true,
		IDX:     3,
	}

	command := provision.GenerateCommand(brr)
	expectedResult := "--bindall true --idx 3"

	if command == "" {
		t.Errorf("incorrect command. got: %s, wanted: %s", command, "")
	}
	if command != expectedResult {
		t.Errorf("incorrect command. got: %s, wanted: %s", command, expectedResult)
	}
}
