package models

type Item struct {
	ID    string
	Price float64
}

type OptimaItems struct {
	OptimaItems   []Item
	SolutionItems []Item
	FinalItems    []string
}
