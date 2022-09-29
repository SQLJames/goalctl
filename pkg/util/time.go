package util

import (
	"fmt"
	"time"

	"github.com/sqljames/goalctl/pkg/log"
)

func StringToTime(inputTime string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, inputTime)
	if err != nil {
		log.Logger.ILog.Error(err, "date is not in the correct format", "inputDate", inputTime)
	}
	return parsedTime, fmt.Errorf("StringToTime: %w", err)
}

func TimeToString(inputTime *time.Time) string {
	return inputTime.UTC().Format(time.RFC3339)
}

