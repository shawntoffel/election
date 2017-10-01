package election

type Counter interface {
	Initialize(config Config)
	Count() (*Result, error)
}

type CounterState struct {
	Events Events
	Error  error
}
