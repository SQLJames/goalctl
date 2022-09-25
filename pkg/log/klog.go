package log

import (
	"flag"
	"os"
	"strconv"

	"github.com/go-logr/logr"
	"github.com/sqljames/goalctl/pkg/info"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
)

type internalklogImplementation struct {
	PanicLogger   logr.Logger
	FatalLogger   logr.Logger
	ErrorLogger   logr.Logger
	WarningLogger logr.Logger
	InfoLogger    logr.Logger
	DebugLogger   logr.Logger
	TraceLogger   logr.Logger
}

func newInternalklog() *internalklogImplementation {
	logger := klogr.NewWithOptions(klogr.WithFormat(klogr.FormatKlog)).
		WithCallDepth(1).
		WithValues("applicationName", info.GetApplicationName())

	return &internalklogImplementation{
		PanicLogger:   logger.WithName("Panic"),
		FatalLogger:   logger.WithName("Fatal"),
		ErrorLogger:   logger.WithName("Error"),
		WarningLogger: logger.WithName("Warn"),
		InfoLogger:    logger.WithName("Info"),
		DebugLogger:   logger.WithName("Debug"),
		TraceLogger:   logger.WithName("Trace"),
	}
}

func (k *internalklogImplementation) Panic(err error, message string, keysAndValues ...interface{}) {
	k.PanicLogger.Error(err, message, keysAndValues...)
	panic(err.Error())
}

func (k *internalklogImplementation) Fatal(err error, message string, keysAndValues ...interface{}) {
	k.FatalLogger.Error(err, message, keysAndValues...)
	os.Exit(1)
}
func (k *internalklogImplementation) Error(err error, message string, keysAndValues ...interface{}) {
	k.ErrorLogger.Error(err, message, keysAndValues...)
}
func (k *internalklogImplementation) Warn(message string, keysAndValues ...interface{}) {
	k.WarningLogger.V(0).Info(message, keysAndValues...)
}

func (k *internalklogImplementation) Info(message string, keysAndValues ...interface{}) {
	k.InfoLogger.V(1).Info(message, keysAndValues...)
}
func (k *internalklogImplementation) Debug(message string, keysAndValues ...interface{}) {
	k.DebugLogger.V(2).Info(message, keysAndValues...)
}
func (k *internalklogImplementation) Trace(message string, keysAndValues ...interface{}) {
	k.TraceLogger.V(3).Info(message, keysAndValues...)
}

// Care of: https://github.com/physcat/klog-cli/blob/main/main.go
func InitializeLogger(logLevel int) {
	fs := flag.NewFlagSet("", flag.PanicOnError)
	klog.InitFlags(fs)
	klog.EnableContextualLogging(true)

	err := fs.Set("v", strconv.Itoa(logLevel))
	if err != nil {
		Logger.Warn("issue setting verbosity flag", "error", err.Error())
	}
}
