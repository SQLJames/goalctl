package tomlprinter

import (
	"fmt"
	"io"

	"github.com/pelletier/go-toml/v2"
	"github.com/sqljames/goalctl/pkg/log"
)

type TomlPrinter struct {
}

func (yp *TomlPrinter) Write(data interface{}, destination io.Writer) (err error) {
	bytes, err := toml.Marshal(data)
	if err != nil {
		return fmt.Errorf("toml: %w", err)
	}

	_, err = destination.Write(bytes)
	if err != nil {
		log.Logger.ILog.Warn("issue writing data out to destination.", "error", err.Error())

		return fmt.Errorf("toml: %w", err)
	}

	return nil
}
