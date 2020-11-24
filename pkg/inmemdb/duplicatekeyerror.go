package inmemdb

import "fmt"

// DuplicateKeyError struct for duplicate key error
type DuplicateKeyError struct {
	key string
}

func (e DuplicateKeyError) Error() string {
	return fmt.Sprintf("Duplicated Key %s", e.key)
}
