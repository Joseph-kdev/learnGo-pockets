package habit

import "fmt"

// InvalidateInputError is returned when user-input data is invalid
type InvalidInputError struct {
	field string
	reason string
}

// Implements error
func(e InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input in field %s: %s", e.field, e.reason)
}