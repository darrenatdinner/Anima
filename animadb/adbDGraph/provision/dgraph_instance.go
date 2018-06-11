package provision

import "projectanima/AnimaUtilities/commandline"

// DGraphInstance - Container For Information To Start A Running Instance Of DGraph
type DGraphInstance struct {
	DirectoryPath  string
	activeInstance bool

	TerminalInstance *commandline.ITerminalInstance
}
