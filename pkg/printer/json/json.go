package jsonprinter

import (
	"encoding/json"
	"io"
)

type JsonPrinter struct {
}

func (jp *JsonPrinter) Write(data interface{}, destination io.Writer) (err error) {
	enc := json.NewEncoder(destination)
	enc.SetIndent("", "  ")
	err = enc.Encode(data)
	if err != nil {
		return err
	}

	return nil
}
