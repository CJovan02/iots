package reading_errors

import "fmt"

var ErrReadingNotFound = &ReadingNotFoundError{}

type ReadingNotFoundError struct {
	Resource string
	Id       uint32
}

func (e *ReadingNotFoundError) Error() string {
	if e.Id == 0 && e.Resource == "" {
		return "Reading not found"
	}

	return fmt.Sprintf("%s with ID %v not found", e.Resource, e.Id)
}

func NewNotFound(resource string, id uint32) error {
	return &ReadingNotFoundError{Resource: resource, Id: id}
}
