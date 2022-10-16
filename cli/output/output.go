package output

import (
	"os"

	"github.com/sqljames/goalctl/pkg/printer"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func Output(OutputFormat string, data interface{}) {
	err := printer.NewPrinterByString(OutputFormat).Writer.Write(data, os.Stdout)
	if err != nil {
		jlogr.Logger.ILog.Warn("issue Printing the data", "function", "ListEntries", "error", err.Error())
	}
}
