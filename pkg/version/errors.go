package version

import "fmt"

type customError struct {
	message string
	err     error
}

func (e *customError) Error() string {
	return fmt.Sprintf("%s: %s", e.message, e.err)
}
