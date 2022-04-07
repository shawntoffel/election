package election

import (
	"io/ioutil"
	"testing"
)

func TestBltExport(t *testing.T) {
	c := generateTestConfig()

	exporter := NewExporter(c)

	expected := loadStringFileData("simple.blt")
	t.Log("expected: \n" + expected)

	actual := exporter.ExportBlt()
	t.Log("actual: \n" + actual)
	if actual != expected {
		t.Error("BLT did not match expected result")
	}
}

func loadStringFileData(filename string) string {
	bytes, _ := ioutil.ReadFile("testdata/" + filename)
	return string(bytes)
}

func generateTestConfig() Config {
	config := Config{
		Title: "Test",
	}

	names := []string{"Alice", "Bob", "Chris", "Don", "Eric", "Frank"}

	for _, name := range names {
		c := Candidate{}
		c.Id = name
		c.Name = name

		config.Candidates = append(config.Candidates, c)
	}

	ballots := []Ballot{}

	ballots = append(ballots, Ballot{
		Count: 28,
		Preferences: []string{
			"Alice",
			"Bob",
			"Chris",
		},
	})

	ballots = append(ballots, Ballot{
		Count: 26,
		Preferences: []string{
			"Bob",
			"Alice",
			"Chris",
		},
	})

	ballots = append(ballots, Ballot{
		Count: 3,
		Preferences: []string{
			"Chris",
		},
	})

	ballots = append(ballots, Ballot{
		Count: 2,
		Preferences: []string{
			"Don",
		},
	})

	ballots = append(ballots, Ballot{
		Count: 1,
		Preferences: []string{
			"Eric",
		},
	})

	config.Ballots = ballots
	config.WithdrawnCandidates = []string{"Frank"}
	config.NumSeats = 3

	return config
}
