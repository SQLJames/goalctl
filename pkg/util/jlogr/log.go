package jlogr

import (
	"github.com/go-logr/logr"
	"github.com/sqljames/goalctl/pkg/info"
	"k8s.io/klog/v2/klogr"
)

type Log struct {
	ILog logger
}

var (
	Logger                    = Log{ILog: newInternalklog()}
	defaultLogger logr.Logger = klogr.NewWithOptions(klogr.WithFormat(klogr.FormatKlog)).
			WithCallDepth(1).
			WithValues("applicationName", info.GetApplicationName())
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
