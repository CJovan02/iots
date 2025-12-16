package sensor

import (
	"fmt"
)

type NotFound struct {
	Id uint32
}

func NewNotFound(id uint32) *NotFound {
	return &NotFound{Id: id}
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("sensor reading #%d not found", e.Id)
}

type InvalidArgument struct {
	Msg string
}

func NewInvalidArgument(msg string) *InvalidArgument {
	return &InvalidArgument{Msg: msg}
}

func (e *InvalidArgument) Error() string {
	return e.Msg
}
