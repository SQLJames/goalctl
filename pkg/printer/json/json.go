package jsonprinter

import (
	"encoding/json"
	"fmt"
	"io"
)

type JSONPrinter struct {
}

func (jp *JSONPrinter) Write(data interface{}, destination io.Writer) (err error) {
	enc := json.NewEncoder(destination)
	enc.SetIndent("", "  ")
	err = enc.Encode(data)
	if err != nil {
		return fmt.Errorf("json: %w", err)
	}

	return nil
}
