package logging

// DatabaseLogger Is An Implementation Of The Logger with connections to a Database
type DatabaseLogger struct {
}

// Log - Logs A New Message
func (dblog DatabaseLogger) Log(logmessage *Message) (DatabaseLogger, error) {
	return dblog, nil
}

// Open - Opens A New Connection
func (dblog DatabaseLogger) Open(log string) (DatabaseLogger, error) {
	return dblog, nil
}

// Close - Closes the Database Connection
func (dblog DatabaseLogger) Close() (bool, error) {
	return false, nil
}

// LoggingType - Method To Get The Type Of Logger
func (dblog DatabaseLogger) LoggingType() Method { return DatabaseLogging }
