package election

import (
	"fmt"
	"io"

	blt "github.com/shawntoffel/goblt"
)

type Config struct {
	Title               string
	NumSeats            int
	Ballots             Ballots
	Candidates          Candidates
	WithdrawnCandidates []string
	Precision           int
	Seed                int64
}

func LoadConfigFromBlt(r io.Reader) (Config, error) {
	parser := blt.NewParser(r)
	result, err := parser.Parse()
	if err != nil {
		return Config{}, fmt.Errorf("failed to parse BLT data: %s", err.Error())
	}

	return Config{
		Title:               result.Title,
		NumSeats:            result.NumSeats,
		Ballots:             createBallots(result),
		Candidates:          createCandidates(result),
		WithdrawnCandidates: result.NamedWithdrawn(),
	}, nil
}

func createCandidates(election *blt.Election) Candidates {
	candidates := make(Candidates, len(election.Candidates))

	for i, name := range election.Candidates {
		candidates[i] = Candidate{
			Id:   name,
			Name: name,
		}
	}

	return candidates
}

func createBallots(election *blt.Election) []Ballot {
	ballots := make([]Ballot, len(election.Ballots))

	for i, parsedBallot := range election.FlatNamedBallots() {
		ballot := Ballot{
			Count:       parsedBallot.Count,
			Preferences: parsedBallot.Preferences,
		}
		ballots[i] = ballot
	}

	return ballots
}
