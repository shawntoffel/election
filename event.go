package election

type Events []Event
type Event struct {
	Description string
}

type ProcessableEvent interface {
	Process() Event
}
