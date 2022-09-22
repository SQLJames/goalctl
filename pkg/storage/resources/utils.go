package resources

import (
	"os/user"
	"time"

	"github.com/sqljames/goalctl/pkg/log"
)

func NewGoal(goal string, dueDate string, details string, priority int) *Goal {
	dueTime, err := time.Parse(time.RFC3339, dueDate)
	if err != nil {
		log.Logger.Error(err, "Unable to get current os user, setting Author to `self`")
		dueTime = time.Now().UTC()
	}
	user, err := user.Current()
	if err != nil {
		log.Logger.Error(err, "Unable to get current os user, setting Author to `self`")
		user.Username = "self"
	}
	return &Goal{
		Author:      user.Username,
		Deadline:    dueTime.Format(time.RFC3339),
		CreatedDate: time.Now().UTC().Format(time.RFC3339),
		Goal:        goal,
		Details:     details,
		Priority:    priority,
		Status:      "Active",
	}
}

func NewLogEntry(entryText string, tags []string) *LogEntry {
	user, err := user.Current()
	if err != nil {
		log.Logger.Error(err, "Unable to get current os user, setting Author to `self`")
		user.Username = "self"
	}
	return &LogEntry{
		Author:      user.Username,
		Tags:        tags,
		CreatedDate: time.Now().UTC().Format(time.RFC3339),
		Entry:       entryText,
	}
}
