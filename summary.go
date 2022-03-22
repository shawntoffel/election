package election

type Summary struct {
	Events Events         `json:"events,omitempty" yaml:"events,omitempty"`
	Rounds RoundSummaries `json:"rounds,omitempty" yaml:"rounds,omitempty"`
}

func (s *Summary) AddEvent(e Event) {
	s.Events = append(s.Events, e)
}

func (s *Summary) AddRound(r RoundSummary) {
	s.Rounds = append(s.Rounds, r)
}

type RoundSummaries []RoundSummary
type RoundSummary struct {
	Number     int                `json:"number" yaml:"number"`
	Excess     int64              `json:"excess" yaml:"excess"`
	Surplus    int64              `json:"surplus" yaml:"surplus"`
	Quota      int64              `json:"quota" yaml:"quota"`
	Candidates []CandidateSummary `json:"candidates" yaml:"candidates"`
}

type CandidateSummary struct {
	Candidate
	Votes  int64  `json:"votes" yaml:"votes"`
	Weight int64  `json:"weight" yaml:"weight"`
	Status string `json:"status" yaml:"status"`
}
