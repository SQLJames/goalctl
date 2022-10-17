package info

const (
	Description string = "A Cli tool for managing your Daily Journal entries."
	Copyright   string = "Database Ally, LLC, 2022"
)

type Author struct {
	Name  string
	Email string
}

var (
	applicationName = "unknown"
	Authors         = []*Author{
		{
			Name:  "James Rhoat",
			Email: "James@Rhoat.com",
		},
	}
)

func GetApplicationName() string {
	return applicationName
}
