package usecases

import (
	"errors"

	"github.com/tavo/prueba/coupon/models"
)

// UseCases entity
type UseCases struct {
	OptimaItems   []models.Item
	SolutionItems []models.Item
	FinalItems    []string
}

// Calculate call to funtion recursive
func (o *UseCases) Calculate(items []models.Item, amount float64) []string {
	o.calculateRec(items, amount, false)
	return o.FinalItems
}

func (o *UseCases) calculateRec(items []models.Item, ammountMaximo float64, optima bool) {
	//validar si es solucion optima
	if optima {
		// comprobar si es mejor solucion que la ultima
		if o.validateOptimalSolution(ammountMaximo) {
			// actualizar solucion
			o.OptimaItems = o.SolutionItems
			o.FinalItems = getIDs(o.SolutionItems)
		}
	} else {
		for _, item := range items {
			// validar si existe en elemento solucion
			if !o.validateExistense(item) {
				// validamos si supera el vamor maximo
				if ammountMaximo > o.CalculateTotal(o.SolutionItems)+item.Price {
					// añadir item a solucion items
					o.SolutionItems = append(o.SolutionItems, item)
					if o.validateOptimalSolution(ammountMaximo) {
						// actualizar UseCases con solucion
						o.calculateRec(items, ammountMaximo, true)
					}
					o.calculateRec(items, ammountMaximo, false)
					// eliminar item a solucion items
					o.SolutionItems = append(o.SolutionItems[:len(o.SolutionItems)-1], o.SolutionItems[len(o.SolutionItems):]...)
				} else {
					o.calculateRec(items, ammountMaximo, true)
				}
			}
		}
	}
}

func (o *UseCases) validateExistense(itemValidate models.Item) bool {
	for _, item := range o.SolutionItems {
		if item.ID == itemValidate.ID {
			return true
		}
	}
	return false
}

func (o *UseCases) validateOptimalSolution(ammountMaximo float64) bool {
	totalSolution := o.CalculateTotal(o.SolutionItems)
	totalOptima := o.CalculateTotal(o.OptimaItems)

	return (totalSolution > totalOptima) && (totalSolution < ammountMaximo)
}

// GetItemWithPrice get price by item
func (o *UseCases) GetItemWithPrice(ids []string, items []models.Item) []models.Item {
	var result []models.Item
	for _, id := range ids {
		for _, item := range items {
			if item.ID == id {
				result = append(result, item)
			}
		}
	}
	return result
}

// ValidatePriceMin get price by item
func (o *UseCases) ValidatePriceMin(price float64, items []models.Item) error {
	var result []models.Item
	for _, item := range items {
		if item.Price < price {
			result = append(result, item)
		}
	}

	if len(result) == 0 {
		return errors.New("ammount insufficiency : ValidatePriceMin")
	}

	return nil
}

// CalculateTotal get total
func (o *UseCases) CalculateTotal(items []models.Item) float64 {
	var sum float64
	for _, item := range items {
		sum = sum + item.Price
	}
	return sum
}

func getIDs(items []models.Item) []string {
	var ids []string
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

//NewUseCases trigger usescases
func NewUseCases(OptimaItems []models.Item, SolutionItems []models.Item, FinalItems []string) *UseCases {
	return &UseCases{OptimaItems, SolutionItems, FinalItems}
}
