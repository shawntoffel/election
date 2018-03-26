package election

type Result struct {
	Candidates Candidates `json:"candidates"`
	Events     Events     `json:"events"`
	Summaries  Events     `json:"summaries"`
}
