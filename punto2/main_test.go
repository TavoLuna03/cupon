package main

import (
	"testing"

	"github.com/tavo/prueba/punto2/models"
)

func TestCalculate(t *testing.T) {
	a := []models.Item{
		{
			ID:    "MLA1",
			Price: 100,
		},
		{
			ID:    "MLA2",
			Price: 210,
		},
		{
			ID:    "MLA3",
			Price: 260,
		},
		{
			ID:    "MLA4",
			Price: 80,
		},
		{
			ID:    "MLA5",
			Price: 90,
		},
	}
	response, _ := calculate(a, 200)
	if len(response) != 100 {
		t.Error("Abs(-1)")
	}
}
