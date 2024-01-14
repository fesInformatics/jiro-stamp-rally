package errors

import "fmt"

type UserDuplicateError struct {
	UserID string
}

func (e UserDuplicateError) Error() string {
	return fmt.Sprintf(`user "%s" is already exists`, e.UserID)
}
