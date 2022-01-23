package election

type Candidates []Candidate
type Candidate struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Rank int    `json:"rank"`
}

type CandidateMap map[string]int

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

func (c Candidates) ToMap() CandidateMap {
	m := CandidateMap{}

	for i := 1; i <= len(c); i++ {
		m[c[i-1].Name] = i
	}

	return m
}
