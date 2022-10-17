package yamlprinter

import (
	"io"

	"github.com/sqljames/goalctl/pkg/util/jlogr"
	"gopkg.in/yaml.v3"
)

type YamlPrinter struct {
}

func (yp *YamlPrinter) Write(data interface{}, destination io.Writer) (err error) {
	bytes, err := yaml.Marshal(data)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return err
	}

	_, err = destination.Write(bytes)
	if err != nil {
		jlogr.Logger.ILog.Error(err, err.Error())

		return err
	}

	return nil
}
