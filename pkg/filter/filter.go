package filter

import (
	"encoding/json"

	"github.com/itchyny/gojq"
	"github.com/sqljames/goalctl/pkg/log"
)

func Filter(filter string, data interface{}) interface{} {
	var marshalTarget interface{}

	query, err := gojq.Parse(filter)
	if err != nil {
		log.Logger.ILog.Error(err, "???")
	}

	marshaledData, err := json.Marshal(data)
	if err != nil {
		log.Logger.ILog.Error(err, "???")
	}

	err = json.Unmarshal(marshaledData, &marshalTarget)
	if err != nil {
		log.Logger.ILog.Error(err, "???")
	}

	iter := query.Run(marshalTarget) // or query.RunWithContext

	for {
		parsedValue, ok := iter.Next()
		if !ok {
			break
		}
		
		if err, ok := parsedValue.(error); ok {
			log.Logger.ILog.Error(err, "???")
		}

		return parsedValue
	}

	return nil
}
