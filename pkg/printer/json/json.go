package jsonprinter

import (
	"encoding/json"
	"io"
)

type JSONPrinter struct {
}

func (jp *JSONPrinter) Write(data interface{}, destination io.Writer) (err error) {
	enc := json.NewEncoder(destination)
	enc.SetIndent("", "  ")
	
	err = enc.Encode(data)
	if err != nil {
		return err
	}

	return nil
}
