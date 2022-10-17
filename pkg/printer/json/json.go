package jsonprinter

import (
	"encoding/json"
	"io"

	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

type JSONPrinter struct {
}

func (jp *JSONPrinter) Write(data interface{}, destination io.Writer) (err error) {
	enc := json.NewEncoder(destination)
	enc.SetIndent("", "  ")

	err = enc.Encode(data)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return err
	}

	return nil
}
