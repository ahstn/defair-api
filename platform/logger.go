package platform

// Fields is intended for passing K/V pairs for structured logging.
type Fields map[string]interface{}

// Logger defines the expected behaviour for a logging implementation.
type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
	WithFields(Fields) Logger
}
