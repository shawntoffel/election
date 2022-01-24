package blt

import (
	"io/ioutil"
	"testing"

	"github.com/shawntoffel/election"
)

func TestBltExport(t *testing.T) {
	c := generateTestConfig()

	exporter := Blt{}

	expected := loadFileData("simple.blt")
	t.Log("expected: \n" + expected)

	actual := exporter.Export(c)
	t.Log("actual: \n" + actual)
	if actual != expected {
		t.Error("BLT did not match expected result")
	}
}

func loadFileData(filename string) string {
	bytes, _ := ioutil.ReadFile("../../testdata/" + filename)
	return string(bytes)
}

func generateTestConfig() election.Config {
	config := election.Config{}

	names := []string{"Alice", "Bob", "Chris", "Don", "Eric", "Frank"}

	for _, name := range names {
		c := election.Candidate{}
		c.Id = name
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
