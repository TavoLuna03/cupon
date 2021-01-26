package usecases

import (
	"reflect"
	"testing"

	"github.com/tavo/prueba/coupon/models"
)

func TestCalculate(t *testing.T) {
	want := []string{"MLA1", "MLA2", "MLA4", "MLA5"}

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

	finalSolution := make([]string, 0)
	usecase := NewUseCases([]models.Item{}, []models.Item{}, finalSolution)
	got := usecase.Calculate(a, float32(500))
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Calculate() = %v, want %v", got, want)
	}
}
