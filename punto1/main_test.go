package main

import (
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	var b OptimaItems
	want := []string{"MLA1", "MLA2", "MLA4", "MLA5"}

	a := []Item{
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
	got := b.calculate(a, 500)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InventoryHttpRepository.PopPickingPosition() = %v, want %v", got, want)
	}
}
