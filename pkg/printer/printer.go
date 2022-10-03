package printer

import (
	"io"
	"strings"

	"github.com/sqljames/goalctl/pkg/flags"
	jsonprinter "github.com/sqljames/goalctl/pkg/printer/json"
	tomlprinter "github.com/sqljames/goalctl/pkg/printer/toml"
	xmlprinter "github.com/sqljames/goalctl/pkg/printer/xml"
	yamlprinter "github.com/sqljames/goalctl/pkg/printer/yaml"
	"github.com/sqljames/goalctl/pkg/util/jlogr"
	"github.com/urfave/cli/v2"
)

var SupportedFormats = []string{"json", "toml", "xml", "yaml"}

type Printer struct {
	Writer Writer
}

type Writer interface {
	Write(data interface{}, writer io.Writer) (err error)
}

func NewPrinter(cliContext *cli.Context) (printer Printer) {
	format := strings.ToLower(cliContext.String(flags.OutputFormatFlagName))

	switch format {
	case SupportedFormats[0]:
		jlogr.Logger.ILog.Trace("Returning json printer")

		return Printer{Writer: &jsonprinter.JSONPrinter{}}
	case SupportedFormats[1]:
		jlogr.Logger.ILog.Trace("Returning toml printer")

		return Printer{Writer: &tomlprinter.TomlPrinter{}}
	case SupportedFormats[2]:
		jlogr.Logger.ILog.Trace("Returning xml printer")

		return Printer{Writer: &xmlprinter.XMLPrinter{}}
	default:
		jlogr.Logger.ILog.Trace("Returning yaml printer")

		return Printer{Writer: &yamlprinter.YamlPrinter{}}
	}
}
