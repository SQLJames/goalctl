package create

import "time"

type NewLogEntry struct {
	LogEntry     string
	NotebookName string
	Tags         []string
}

type NewGoal struct {
	Goal     string
	DueDate  *time.Time
	Details  string
	Priority int
}
