package blt

import (
	"fmt"
	"strings"

	"github.com/shawntoffel/election"
)

var _ election.Exporter = &Blt{}

type Blt struct {
}

func (b Blt) Export(config election.Config) string {
	candidateMap := config.Candidates.ToMap()

	sb := strings.Builder{}

	sb.WriteString(b.buildHeaderLine(len(config.Candidates), config.NumSeats))
	sb.WriteString(b.buildWithdrawnCandidatesLine(config.WithdrawnCandidates, candidateMap))
	sb.WriteString(b.buildBallotLines(config.Ballots.Rollup(), candidateMap))
	sb.WriteString(b.buildCandidateLines(candidateMap))

	return sb.String()
}

func (b Blt) buildHeaderLine(numCandidates int, numSeats int) string {
	return fmt.Sprintf("%d %d \n", numCandidates, numSeats)
}

func (b Blt) buildWithdrawnCandidatesLine(withdrawnCandidates []string, candidateMap election.CandidateMap) string {
	if len(withdrawnCandidates) < 1 {
		return ""
	}

	sb := strings.Builder{}

	for _, wc := range withdrawnCandidates {
		c := candidateMap[wc]
		fmt.Fprintf(&sb, "-%d ", c)
	}

	sb.WriteString("\n")

	return sb.String()
}

func (b Blt) buildBallotLines(rolledUpBallots election.RolledUpBallots, candidateMap election.CandidateMap) string {
	lines := BltBallotLines{}

	for _, rolledUpBallot := range rolledUpBallots {
		line := NewBltBallotLine(rolledUpBallot, candidateMap)
		lines = append(lines, line)
	}

	return lines.String()
}

func (b Blt) buildCandidateLines(candidateMap election.CandidateMap) string {
	bltCandidates := BltCandidates{}

	for n, v := range candidateMap {
		bltCandidate := BltCandidate{
			Id:   v,
			Name: n,
		}

		bltCandidates = append(bltCandidates, bltCandidate)
	}

	return bltCandidates.String()
}
