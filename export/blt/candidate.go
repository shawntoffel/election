package blt

import (
	"fmt"
	"sort"
	"strings"

	"github.com/shawntoffel/election"
)

type BltCandidates []BltCandidate
type BltCandidate struct {
	Id        int
	Candidate election.Candidate
}

type BltCandidateMap map[string]BltCandidate

func (b BltCandidate) String() string {
	return fmt.Sprintf("%d. %s", b.Id, b.Candidate.Name)
}

func (b BltCandidates) String() string {
	sb := strings.Builder{}

	sort.Slice(b, func(i, j int) bool {
		return b[i].Id < b[j].Id
	})

	for _, candidate := range b {
		sb.WriteString(candidate.String() + "\r\n")
	}

	return sb.String()
}
