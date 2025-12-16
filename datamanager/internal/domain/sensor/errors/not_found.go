package sensor

import "fmt"

type NotFound struct {
	Id uint32
}

func NewNotFound(id uint32) *NotFound {
	return &NotFound{Id: id}
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("sensor reading #%d not found", e.Id)
}
