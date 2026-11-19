package repository

// Error is used to define sentinel errors
type Error string

// Error implements the error interface
func (r Error) Error() string {
	return string(r)
}

const (
	ErrNotFound = Error("habit not found in repository")
	ErrConflictingID = Error("cannot create a habit with already existing ID")
)