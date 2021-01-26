package models

type Item struct {
	ID    string
	Price float32
}

type OptimaItems struct {
	OptimaItems   []Item
	SolutionItems []Item
	FinalItems    []string
}
