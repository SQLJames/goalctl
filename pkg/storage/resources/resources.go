package resources

type Book struct {
	Journal Journal `json:"journal,omitempty" toml:"journal,omitempty" xml:"journal,omitempty" yaml:"journal,omitempty"`
}
type Journal struct {
	NoteBooks   []Notebook   `json:"notebooks,omitempty" toml:"notebooks,omitempty" xml:"notebooks,omitempty" yaml:"notebooks,omitempty"`
	GoalDetails []GoalDetail `json:"goalDetails,omitempty" toml:"goalDetails,omitempty" xml:"goalDetails,omitempty" yaml:"goalDetails,omitempty"`
}

type GoalDetail struct {
	Goal    Goal       `json:"goalEntry,omitempty"  toml:"goalEntry,omitempty"  xml:"goalEntry,omitempty"  yaml:"goalEntry,omitempty"`
	Entries []LogEntry `json:"logEntries,omitempty"  toml:"logEntries,omitempty"  xml:"logEntries,omitempty"  yaml:"logEntries,omitempty"`
}

type Goal struct {
	// GoalID Primary key of the goal struct
	GoalID int `json:"goalID,omitempty"  toml:"goalID,omitempty"  xml:"goalID,omitempty"  yaml:"goalID,omitempty"`
	// Author of the goal
	Author string `json:"author,omitempty"  toml:"author,omitempty"  xml:"author,omitempty"  yaml:"author,omitempty"`
	// Deadline for the goal time format RFC3339
	Deadline string `json:"deadline,omitempty"  toml:"deadline,omitempty"  xml:"deadline,omitempty"  yaml:"deadline,omitempty"`
	// CreatedDate is the date the goal was created time format RFC3339
	CreatedDate string `json:"createdDate,omitempty"  toml:"createdDate,omitempty"  xml:"createdDate,omitempty"  yaml:"createdDate,omitempty"`
	// Goal is what you would like to accomplish
	Goal string `json:"goal,omitempty"  toml:"goal,omitempty"  xml:"goal,omitempty"  yaml:"goal,omitempty"`
	// Details of the goal
	Details string `json:"details,omitempty"  toml:"details,omitempty"  xml:"details,omitempty"  yaml:"details,omitempty"`
	// Priority of the goal
	Priority int `json:"priority,omitempty"  toml:"priority,omitempty"  xml:"priority,omitempty"  yaml:"priority,omitempty"`
	// Priority of the goal
	Status string `json:"status,omitempty"  toml:"status,omitempty"  xml:"status,omitempty"  yaml:"status,omitempty"`
}

type Association struct {
	GoalID     int
	LogEntryID int
}

type Notebook struct {
	Notebookid int64      `json:"notebookid,omitempty"        toml:"notebookid,omitempty"        xml:"notebookid,omitempty"        yaml:"notebookid,omitempty"`
	Name       string     `json:"notebookName,omitempty" toml:"notebookName,omitempty" xml:"notebookName,omitempty" yaml:"notebookName,omitempty"`
	Entries    []LogEntry `json:"entries,omitempty"      toml:"entries,omitempty"      xml:"entries,omitempty"      yaml:"entries,omitempty"`
}

type LogEntry struct {
	LogEntryID  int64    `json:"logentryid,omitempty"        toml:"logentryid,omitempty"        xml:"logentryid,omitempty"        yaml:"logentryid,omitempty"`
	Author      string   `json:"author,omitempty"        toml:"author,omitempty"        xml:"author,omitempty"        yaml:"author,omitempty"`
	Tags        []string `json:"tags,omitempty"           toml:"tags,omitempty"           xml:"tags,omitempty"           yaml:"tags,omitempty"`
	Entry       string   `json:"entry,omitempty"           toml:"entry,omitempty"           xml:"entry,omitempty"           yaml:"entry,omitempty"`
	CreatedDate string   `json:"createdDate,omitempty"   toml:"createdDate,omitempty"   xml:"createdDate,omitempty"   yaml:"createdDate,omitempty"`
	Notebookid  int64    `json:"notebookid,omitempty"        toml:"notebookid,omitempty"        xml:"notebookid,omitempty"        yaml:"notebookid,omitempty"`
}
