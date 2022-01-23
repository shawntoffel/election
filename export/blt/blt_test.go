package blt

import (
	"testing"

	"github.com/shawntoffel/election"
)

func TestBltExport(t *testing.T) {
	c := generateTestConfig()

	exporter := Blt{}

	actual := exporter.Export(c)

	t.Log("actual: \n" + actual)

	expected := `6 3 
-6 
28 1 2 3 0
26 2 1 3 0
3 3 0
2 4 0
1 5 0
0
Alice
Bob
Chris
Don
Eric
Frank
`
	t.Log("expected: \n" + expected)

	if actual != expected {
		t.Error("BLT did not match expected result")
	}
}

func generateTestConfig() election.Config {
	config := election.Config{}

	names := []string{"Alice", "Bob", "Chris", "Don", "Eric", "Frank"}

	for _, name := range names {
		c := election.Candidate{}
		c.Name = name

		config.Candidates = append(config.Candidates, c)
	}

	var ballots election.Ballots

	for i := 0; i < 28; i++ {
		var ballot = election.NewBallot()
		ballot.PushBack("Alice")
		ballot.PushBack("Bob")
		ballot.PushBack("Chris")
		ballots = append(ballots, ballot)
	}

	for i := 0; i < 26; i++ {
		var ballot = election.NewBallot()
		ballot.PushBack("Bob")
		ballot.PushBack("Alice")
		ballot.PushBack("Chris")
		ballots = append(ballots, ballot)
	}

	for i := 0; i < 3; i++ {
		var ballot = election.NewBallot()
		ballot.PushBack("Chris")
		ballots = append(ballots, ballot)
	}

	for i := 0; i < 2; i++ {
		var ballot = election.NewBallot()
		ballot.PushBack("Don")
		ballots = append(ballots, ballot)
	}

	var ballot = election.NewBallot()
	ballot.PushBack("Eric")
	ballots = append(ballots, ballot)

	config.Ballots = ballots
	config.WithdrawnCandidates = []string{"Frank"}
	config.NumSeats = 3

	return config
}
