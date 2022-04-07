package election

import (
	"fmt"
	"sort"
	"strings"
)

type Exporter struct {
	config Config
}

type BltBallotLines []BltBallotLine
type BltBallotLine struct {
	Count   int
	Content string
}

type BltCandidates []BltCandidate
type BltCandidate struct {
	Id        int
	Candidate Candidate
}
type BltCandidateMap map[string]BltCandidate

func (b BltCandidate) String() string {
	return fmt.Sprintf("\"%d. %s\"", b.Id, b.Candidate.Name)
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

func NewExporter(config Config) *Exporter {
	return &Exporter{
		config: config,
	}
}

func (e *Exporter) ExportBlt() string {
	candidateMap := e.buildCandidateMap()

	sb := strings.Builder{}

	sb.WriteString(e.buildHeaderLine())
	sb.WriteString(e.buildWithdrawnCandidatesLine(candidateMap))
	sb.WriteString(e.buildBallotLines(candidateMap))
	sb.WriteString(e.buildCandidateLines(candidateMap))
	sb.WriteString(e.buildTitle())

	return sb.String()
}

func (e *Exporter) buildHeaderLine() string {
	return fmt.Sprintf("%d %d\r\n", len(e.config.Candidates), e.config.NumSeats)
}

func (e *Exporter) buildWithdrawnCandidatesLine(candidateMap BltCandidateMap) string {
	withdrawnCandidates := e.config.WithdrawnCandidates
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

func (e *Exporter) buildBallotLines(candidateMap BltCandidateMap) string {
	lines := BltBallotLines{}

	for _, ballot := range e.config.Ballots {
		line := newBltBallotLine(ballot, candidateMap)
		lines = append(lines, line)
	}

	return lines.String()
}

func (e *Exporter) buildCandidateLines(candidateMap BltCandidateMap) string {
	bltCandidates := BltCandidates{}

	for _, v := range candidateMap {
		bltCandidates = append(bltCandidates, v)
	}

	return bltCandidates.String()
}

func (e *Exporter) buildTitle() string {
	text := e.config.Title
	if text == "" {
		text = "Election"
	}

	return "\"" + text + "\""
}

func (e *Exporter) buildCandidateMap() BltCandidateMap {
	candidates := e.config.Candidates
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

func newBltBallotLine(ballot Ballot, candidateMap BltCandidateMap) BltBallotLine {
	sb := strings.Builder{}

	for _, pref := range ballot.Preferences {
		c := candidateMap[pref]
		fmt.Fprintf(&sb, "%d ", c.Id)
	}

	sb.WriteString("0")

	return BltBallotLine{
		Count:   ballot.Count,
		Content: sb.String(),
	}
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
		fmt.Fprintf(&sb, line.String()+"\r\n")
	}

	sb.WriteString("0\r\n")

	return sb.String()
}
