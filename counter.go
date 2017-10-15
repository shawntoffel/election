package election

type Counter interface {
	Initialize(config Config) error
	Count() (*Result, error)
}

type CounterState struct {
	Events Events
	Error  error
}

func (s *CounterState) AddEvent(event ProcessableEvent) {
	e := event.Process()

	s.Events = append(s.Events, e)
}
