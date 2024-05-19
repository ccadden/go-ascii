package cmderrors

import "fmt"

type CharError struct {
	Err         error
	Message     string
	NumAsString string
}

func (e *CharError) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.NumAsString)
}
