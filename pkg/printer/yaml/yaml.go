package yamlprinter

import (
	"fmt"
	"io"

	"github.com/sqljames/goalctl/pkg/log"

	"gopkg.in/yaml.v3"
)

type YamlPrinter struct {
}

func (yp *YamlPrinter) Write(data interface{}, destination io.Writer) (err error) {
	bytes, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("yaml: %w", err)
	}

	_, err = destination.Write(bytes)
	if err != nil {
		log.Logger.ILog.Warn("issue writing data out to destination.", "error", err.Error())
	}

	return nil
}
