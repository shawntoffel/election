package election

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestLoadFromBlt(t *testing.T) {
	data := loadFileData("simple.blt")
	c, err := LoadConfigFromBlt(data)
	if err != nil {
		t.Errorf("failed to load config from blt: %s", err.Error())
	}

	assertInt(t, 3, c.NumSeats, "seats")
	assertInt(t, 1, len(c.WithdrawnCandidates), "withdrawn")
	assertInt(t, 6, len(c.Candidates), "candidates")
	assertInt(t, 5, len(c.Ballots), "ballots")

	sum := 0
	for _, ballot := range c.Ballots {
		sum += len(ballot.Preferences)
	}

	assertInt(t, 9, sum, "preferences")
}

func assertInt(t *testing.T, expected, got int, name string) {
	if expected != got {
		t.Error("expected", expected, name, ". got", got)
	}
}

func loadFileData(filename string) io.Reader {
	bytes, _ := ioutil.ReadFile("testdata/" + filename)
	return strings.NewReader(string(bytes))
}
