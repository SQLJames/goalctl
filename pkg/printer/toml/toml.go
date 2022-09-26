package tomlprinter

import (
	"io"

	"github.com/pelletier/go-toml/v2"
	"github.com/sqljames/goalctl/pkg/log"
)

type TomlPrinter struct {
}

func (yp *TomlPrinter) Write(data interface{}, destination io.Writer) (err error) {
	bytes, err := toml.Marshal(data)
	if err != nil {
		return err
	}
	
	_, err = destination.Write(bytes)
	if err != nil {
		log.Logger.Warn("issue writing data out to destination.", "error", err.Error())
	}

	return nil
}
