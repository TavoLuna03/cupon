package models

type Item struct {
	ID    string
	Price int
}

type OptimaItems struct {
	OptimaItems   []Item
	SolutionItems []Item
	FinalItems    []string
}
