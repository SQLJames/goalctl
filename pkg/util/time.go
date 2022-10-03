package util

import (
	"fmt"
	"time"

	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func StringToTime(inputTime string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, inputTime)
	if err != nil {
		jlogr.Logger.ILog.Error(err, "date is not in the correct format", "inputDate", inputTime)
		err = fmt.Errorf("StringToTime: %w", err)
	}

	return parsedTime, nil
}

func TimeToString(inputTime *time.Time) string {
	return inputTime.UTC().Format(time.RFC3339)
}
