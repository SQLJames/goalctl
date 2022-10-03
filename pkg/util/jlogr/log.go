package jlogr

type Log struct {
	ILog logger
}

var (
	Logger = Log{ILog: newInternalklog()}
)

type logger interface {
	Panic(err error, message string, keysAndValues ...interface{})
	Fatal(err error, message string, keysAndValues ...interface{})
	Error(err error, message string, keysAndValues ...interface{})
	Warn(message string, keysAndValues ...interface{})
	Info(message string, keysAndValues ...interface{})
	Debug(message string, keysAndValues ...interface{})
	Trace(message string, keysAndValues ...interface{})
}
