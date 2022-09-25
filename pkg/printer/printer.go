package printer

import (
	"io"
	"strings"

	"github.com/sqljames/goalctl/pkg/flags"
	"github.com/sqljames/goalctl/pkg/log"
	jsonprinter "github.com/sqljames/goalctl/pkg/printer/json"
	tomlprinter "github.com/sqljames/goalctl/pkg/printer/toml"
	xmlprinter "github.com/sqljames/goalctl/pkg/printer/xml"
	yamlprinter "github.com/sqljames/goalctl/pkg/printer/yaml"
	"github.com/urfave/cli/v2"
)

var SupportedFormats = []string{"json", "toml", "xml", "yaml"}

type Printer interface {
	Write(data interface{}, writer io.Writer) (err error)
}

func NewPrinter(cliContext *cli.Context) (printer Printer) {
	format := cliContext.String(flags.OutputFormatFlagName)
	format = strings.ToLower(format)
	if format == "" || !contains(SupportedFormats, format) {
		log.Logger.Trace("format not defined or is unsupported, using default format.")
	}
	switch format {
	case SupportedFormats[0]:
		log.Logger.Trace("Returning json printer")
		return &jsonprinter.JSONPrinter{}
	case SupportedFormats[1]:
		log.Logger.Trace("Returning toml printer")
		return &tomlprinter.TomlPrinter{}
	case SupportedFormats[2]:
		log.Logger.Trace("Returning xml printer")
		return &xmlprinter.XMLPrinter{}
	default:
		log.Logger.Trace("Returning yaml printer")
		return &yamlprinter.YamlPrinter{}
	}
}

func contains(list []string, value string) (supportedFormat bool) {
	for _, listitem := range list {
		if value == listitem {
			return true
		}
	}
	return false
}
