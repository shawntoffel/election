package election

type Result struct {
	Title      string     `json:"title,omitempty" yaml:"title,omitempty"`
	NumSeats   int        `json:"numSeats,omitempty" yaml:"numSeats,omitempty"`
	NumBallots int        `json:"numBallots,omitempty" yaml:"numBallots,omitempty"`
	Precision  int        `json:"precision,omitempty" yaml:"precision,omitempty"`
	Seed       int64      `json:"seed,omitempty" yaml:"seed,omitempty"`
	Elected    Candidates `json:"elected,omitempty" yaml:"elected,omitempty"`
	Summary    *Summary   `json:"summary,omitempty" yaml:"summary,omitempty"`
}
