package logging

// Method Indicates The Current Method Of Logging Instanciated
type Method int

const (
	// FileLogging Constant
	FileLogging Method = iota
	// DatabaseLogging Constant
	DatabaseLogging Method = iota
)

// ILogger Interface For Creating And Modifying Error/Data Logs
type ILogger interface {
	Log(*Message) (ILogger, error)
	Open(string) (ILogger, error)
	Close() (bool, error)
	LoggingType() int
}

// GetLogger - Create A Logger Of The Given Type
func GetLogger(loggerType Method) *ILogger {
	switch loggerType {
	case FileLogging:

		return nil
	case DatabaseLogging:

		return nil
	}
	return nil
}
