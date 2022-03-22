package election

type Counter interface {
	Initialize(config Config) error
	Count() (*Result, error)
}

type CounterState struct {
	Summary *Summary
	Error   error
}

func (s *CounterState) AddEvent(e ProcessableEvent) {
	if s.Summary == nil {
		s.Summary = &Summary{}
	}

	s.Summary.AddEvent(NewEvent(e))
}
