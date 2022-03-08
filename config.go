package election

type Config struct {
	Title               string
	NumSeats            int
	Ballots             Ballots
	Candidates          Candidates
	WithdrawnCandidates []string
	Precision           int
	Seed                int64
}
