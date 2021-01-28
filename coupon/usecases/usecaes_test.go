package usecases

import (
	"errors"
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
	got := usecase.Calculate(a, float32(500.00))
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Calculate() = %v, want %v", got, want)
	}
}

func TestCalculateSuperate(t *testing.T) {
	want := []string{"MLA1", "MLA2", "MLA3", "MLA4", "MLA5"}

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
	got := usecase.Calculate(a, float32(1000))
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Calculate() = %v, want %v", got, want)
	}
}

func TestCalculateValidatePriceFail(t *testing.T) {
	want := errors.New("ammount insufficiency : ValidatePriceMin")

	a := []models.Item{
		{
			ID:    "MLA1",
			Price: 100,
		},
	}

	finalSolution := make([]string, 0)
	usecase := NewUseCases([]models.Item{}, []models.Item{}, finalSolution)
	err := usecase.ValidatePriceMin(10, a)
	if !reflect.DeepEqual(err, want) {
		t.Errorf("Calculate() = %v, want %v", err, want)
	}
}

func TestGetIds(t *testing.T) {
	want := []models.Item{
		{
			ID:    "MLA1",
			Price: 100,
		},
	}

	a := []string{"MLA1"}

	items := []models.Item{
		{
			ID:    "MLA1",
			Price: 100,
		},
	}

	finalSolution := make([]string, 0)
	usecase := NewUseCases([]models.Item{}, []models.Item{}, finalSolution)
	got := usecase.GetItemWithPrice(a, items)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Calculate() = %v, want %v", got, want)
	}
}
