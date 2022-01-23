package blt

import (
	"sort"
	"strings"
)

type BltCandidates []BltCandidate
type BltCandidate struct {
	Id   int
	Name string
}

func (b BltCandidate) String() string {
	return b.Name
}

func (b BltCandidates) String() string {
	sb := strings.Builder{}

	sort.Slice(b, func(i, j int) bool {
		return b[i].Id < b[j].Id
	})

	for _, candidate := range b {
		sb.WriteString(candidate.String() + "\n")
	}

	return sb.String()
}
