package election

type Candidates []Candidate

func (c Candidates) Equals(compare Candidates) bool {
	if len(c) != len(compare) {
		return false
	}
	for i := range c {
		if c[i] != compare[i] {
			return false
		}
	}
	return true
}

type Candidate struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Rank int    `json:"rank"`
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
