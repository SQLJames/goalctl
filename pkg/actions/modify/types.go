package modify

import "time"

type GoalModificationOptions struct {
	GoalDeadline *time.Time
	GoalName     string
	GoalDetails  string
	GoalPriority int
	GoalStatus   string
}

type EntryModificationOptions struct {
	TargetNotebookID int64
	EntryTags        []string
	EntryDetails     string
}
