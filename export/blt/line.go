package blt

import (
	"fmt"
	"sort"
	"strings"

	"github.com/shawntoffel/election"
)

type BltBallotLines []BltBallotLine
type BltBallotLine struct {
	Count   int
	Content string
}

func NewBltBallotLine(rolledUpBallot election.RolledUpBallot, candidateMap election.CandidateMap) BltBallotLine {
	sb := strings.Builder{}

	line := BltBallotLine{
		Count: rolledUpBallot.Count,
	}

	iter := rolledUpBallot.Ballot.List.Front()
	for {
		c := candidateMap[iter.Value.(string)]

		fmt.Fprintf(&sb, "%d ", c)

		if iter.Next() == nil {
			break
		}

		iter = iter.Next()
	}

	sb.WriteString("0")

	line.Content = sb.String()

	return line
}

func (b BltBallotLine) String() string {
	return fmt.Sprintf("%d %s", b.Count, b.Content)
}

func (b BltBallotLines) String() string {
	sb := strings.Builder{}

	sort.Slice(b, func(i, j int) bool {
		return b[i].Count > b[j].Count
	})

	for _, line := range b {
		fmt.Fprintf(&sb, line.String()+"\n")
	}

	sb.WriteString("0\n")

	return sb.String()
}
