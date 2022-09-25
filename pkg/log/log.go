package log

var (
	Logger = newLogger()
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

func newLogger() logger {
	return newInternalklog()
}
