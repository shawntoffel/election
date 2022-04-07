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
	assertStringSlice(t, []string{"6. Frank"}, c.WithdrawnCandidates, "withdrawn")
	assertInt(t, 6, len(c.Candidates), "candidates")
	assertInt(t, 5, len(c.Ballots), "ballots")

	ballots := c.Ballots
	assertInt(t, ballots[0].Count, 28, "ballots")
	assertStringSlice(t, []string{
		"1. Alice",
		"2. Bob",
		"3. Chris",
	}, ballots[0].Preferences, "preferences")

	assertInt(t, ballots[1].Count, 26, "ballots")
	assertStringSlice(t, []string{
		"2. Bob",
		"1. Alice",
		"3. Chris",
	}, ballots[1].Preferences, "preferences")

	assertInt(t, ballots[2].Count, 3, "ballots")
	assertStringSlice(t, []string{"3. Chris"}, ballots[2].Preferences, "preferences")

	assertInt(t, ballots[3].Count, 2, "ballots")
	assertStringSlice(t, []string{"4. Don"}, ballots[3].Preferences, "preferences")

	assertInt(t, ballots[4].Count, 1, "ballots")
	assertStringSlice(t, []string{"5. Eric"}, ballots[4].Preferences, "preferences")
}

func assertInt(t *testing.T, expected, got int, name string) {
	if expected != got {
		t.Error("expected", expected, name, ": got", got)
	}
}

func assertStringSlice(t *testing.T, expected, got []string, name string) {
	if !testEqualStringSlice(expected, got) {
		t.Error("expected", expected, name, ": got", got)
	}
}

func testEqualStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func loadFileData(filename string) io.Reader {
	bytes, _ := ioutil.ReadFile("testdata/" + filename)
	return strings.NewReader(string(bytes))
}
