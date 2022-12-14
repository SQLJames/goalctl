package resources

import (
	"os/user"
	"time"

	"github.com/sqljames/goalctl/pkg/util/jlogr"
)

func NewGoal(goal, dueDate, details string, priority int) *Goal {
	dueTime, err := time.Parse(time.RFC3339, dueDate)
	if err != nil {
		jlogr.Logger.ILog.Error(err, "Unable to parse due time , getting current time.")

		dueTime = time.Now().UTC()
	}

	return &Goal{
		Author:      getCurrentUser(),
		Deadline:    dueTime.Format(time.RFC3339),
		CreatedDate: time.Now().UTC().Format(time.RFC3339),
		Goal:        goal,
		Details:     details,
		Priority:    priority,
		Status:      "Active",
	}
}

func NewLogEntry(entryText string, tags []string) *LogEntry {
	return &LogEntry{
		Author:      getCurrentUser(),
		Tags:        tags,
		CreatedDate: time.Now().UTC().Format(time.RFC3339),
		Entry:       entryText,
	}
}

func getCurrentUser() (username string) {
	currentUser, err := user.Current()
	if err != nil {
		jlogr.Logger.ILog.Error(err, "Unable to get current os user, setting Author to `self`")

		currentUser.Username = "self"
	}

	return currentUser.Username
}
