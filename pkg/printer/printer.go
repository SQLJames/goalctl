package printer

import (
	"io"
	"strings"

	jsonprinter "github.com/sqljames/goalctl/pkg/printer/json"
	tomlprinter "github.com/sqljames/goalctl/pkg/printer/toml"
	xmlprinter "github.com/sqljames/goalctl/pkg/printer/xml"
	yamlprinter "github.com/sqljames/goalctl/pkg/printer/yaml"
)

type Type int8

const (
	JSONPrinter Type = iota
	TOMLPrinter
	XMLPrinter
	YAMLPrinter
)

type Printer struct {
	Writer Writer
}

type Writer interface {
	Write(data interface{}, writer io.Writer) (err error)
}

func NewPrinterByType(printerType Type) (printer Printer) {
	switch printerType {
	case JSONPrinter:
		return Printer{Writer: &jsonprinter.JSONPrinter{}}
	case TOMLPrinter:
		return Printer{Writer: &tomlprinter.TomlPrinter{}}
	case XMLPrinter:
		return Printer{Writer: &xmlprinter.XMLPrinter{}}
	case YAMLPrinter:
		return Printer{Writer: &yamlprinter.YamlPrinter{}}
	default:
		return Printer{Writer: &yamlprinter.YamlPrinter{}}
	}
}

func NewPrinterByString(outputFormat string) (printer Printer) {
	format := strings.ToLower(outputFormat)
	switch format {
	case "json":
		return Printer{Writer: &jsonprinter.JSONPrinter{}}
	case "toml":
		return Printer{Writer: &tomlprinter.TomlPrinter{}}
	case "xml":
		return Printer{Writer: &xmlprinter.XMLPrinter{}}
	case "yaml":
		return Printer{Writer: &yamlprinter.YamlPrinter{}}
	default:
		return Printer{Writer: &yamlprinter.YamlPrinter{}}
	}
}
