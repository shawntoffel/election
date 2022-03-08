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
	candidateMap := b.buildCandidateMap(config.Candidates)

	sb := strings.Builder{}

	sb.WriteString(b.buildHeaderLine(len(config.Candidates), config.NumSeats))
	sb.WriteString(b.buildWithdrawnCandidatesLine(config.WithdrawnCandidates, candidateMap))
	sb.WriteString(b.buildBallotLines(config.Ballots.Rollup(), candidateMap))
	sb.WriteString(b.buildCandidateLines(candidateMap))
	sb.WriteString(b.buildTitle(config.Title))

	return sb.String()
}

func (b Blt) buildHeaderLine(numCandidates int, numSeats int) string {
	return fmt.Sprintf("%d %d\r\n", numCandidates, numSeats)
}

func (b Blt) buildWithdrawnCandidatesLine(withdrawnCandidates []string, candidateMap BltCandidateMap) string {
	if len(withdrawnCandidates) < 1 {
		return ""
	}

	sb := strings.Builder{}

	for _, wc := range withdrawnCandidates {
		for _, c := range candidateMap {
			if strings.EqualFold(c.Candidate.Name, wc) {
				fmt.Fprintf(&sb, "-%d ", c.Id)
			}
		}
	}

	sb.WriteString("\r\n")

	return sb.String()
}

func (b Blt) buildBallotLines(rolledUpBallots election.RolledUpBallots, candidateMap BltCandidateMap) string {
	lines := BltBallotLines{}

	for _, rolledUpBallot := range rolledUpBallots {
		line := NewBltBallotLine(rolledUpBallot, candidateMap)
		lines = append(lines, line)
	}

	return lines.String()
}

func (b Blt) buildCandidateLines(candidateMap BltCandidateMap) string {
	bltCandidates := BltCandidates{}

	for _, v := range candidateMap {
		bltCandidates = append(bltCandidates, v)
	}

	return bltCandidates.String()
}

func (b Blt) buildTitle(title string) string {
	text := title
	if title == "" {
		text = "Election"
	}

	return "\"" + text + "\""
}

func (b Blt) buildCandidateMap(candidates election.Candidates) BltCandidateMap {
	m := BltCandidateMap{}

	for i := 1; i <= len(candidates); i++ {

		candidate := candidates[i-1]

		bltCandidate := BltCandidate{
			Id:        i,
			Candidate: candidate,
		}

		m[candidate.Id] = bltCandidate
	}

	return m
}
