package election

type Candidates []Candidate
type Candidate struct {
	Id   string
	Name string
	Rank int
}

type ByRank Candidates

func (c ByRank) Len() int {
	return len(c)
}

func (c ByRank) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c ByRank) Less(i, j int) bool {
	return c[i].Rank < c[j].Rank
}
