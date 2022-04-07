package election

type Ballot struct {
	Count       int
	Preferences []string
}

type Ballots []Ballot

func (b Ballots) TotalCount() int {
	total := 0
	for _, ballot := range b {
		total += ballot.Count
	}
	return total
}
