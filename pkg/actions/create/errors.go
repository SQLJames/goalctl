package create

type EmptryStringError struct{}

func (empty *EmptryStringError) Error() string {
	return "string can not be empty"
}
