package election

import (
	"reflect"
)

type Events []Event
type Event struct {
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

func NewEvent(e ProcessableEvent) Event {
	return Event{
		Type:        reflect.TypeOf(e).Elem().Name(),
		Description: e.Process(),
	}
}

type ProcessableEvent interface {
	Process() string
}
