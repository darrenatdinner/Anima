package provision

import "projectanima/AnimaUtilities/commandline"

// DGraphZeroInstance is a container for accessing DGraph's underlying data controller
type DGraphZeroInstance struct {
	DirectoryPath  string
	activeInstance bool

	TerminalInstance *commandline.ITerminalInstance
}
