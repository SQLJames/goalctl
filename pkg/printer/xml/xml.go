package xmlprinter

import (
	"encoding/xml"
	"io"

	"github.com/sqljames/goalctl/pkg/log"
)

type XMLPrinter struct {
}

func (yp *XMLPrinter) Write(data interface{}, destination io.Writer) (err error) {
	enc := xml.NewEncoder(destination)
	enc.Indent("", "  ")
	err = enc.Encode(data)
	if err != nil {
		return err
	}
	// the xml library doesnt encode a newline into the marshaller like the other libraries
	// in ZSH this can result in a percent sign (%) being placed at the end of the data.

	_, err = destination.Write([]byte("\n"))
	if err != nil {
		log.Logger.Warn("issue writing data out to destination.", "error", err.Error())
	}
	
	return err
}
