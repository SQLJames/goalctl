package actions

type NoModificationError struct{}

func (m *NoModificationError) Error() string {
	return "no values provided for modification."
}