package output

import (
	"os"

	"github.com/sqljames/goalctl/pkg/printer"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func Output(outputFormat string, data interface{}) {
	err := printer.NewPrinterByString(outputFormat).Writer.Write(data, os.Stdout)
	if err != nil {
		jlogr.Logger.ILog.Error(err, "issue Printing the data", "function", "ListEntries", "error", err.Error())
	}
}
