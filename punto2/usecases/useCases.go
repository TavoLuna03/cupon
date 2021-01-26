package usecases

import (
	"fmt"

	"github.com/tavo/prueba/punto2/models"
)

type UseCases struct {
	OptimaItems   []models.Item
	SolutionItems []models.Item
	FinalItems    []string
}

func (o *UseCases) Calculate(items []models.Item, amount int) []string {
	o.calculateRec(items, amount, false)
	return o.FinalItems
}

func (o *UseCases) calculateRec(items []models.Item, ammountMaximo int, optima bool) {
	//validar si es solucion optima
	if optima {
		// comprobar si es mejor solucion que la ultima
		if o.validateOptimalSolution(ammountMaximo) {
			// actualizar UseCasescon solucion
			fmt.Printf("SolutionItems")
			fmt.Println(o.SolutionItems)
			o.OptimaItems = o.SolutionItems
			o.FinalItems = getIDs(o.SolutionItems)
		}
	} else {
		for _, item := range items {
			// validar si existe en elemento solucion
			if !o.validateExistense(item) {
				// validamos si supera el vamor maximo
				if ammountMaximo > calculateTotal(o.SolutionItems)+item.Price {
					// añadir item a solucion items
					o.SolutionItems = append(o.SolutionItems, item)
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

func (o *UseCases) validateOptimalSolution(ammountMaximo int) bool {
	totalSolution := calculateTotal(o.SolutionItems)
	totalOptima := calculateTotal(o.OptimaItems)

	return (totalSolution > totalOptima) && (totalSolution < ammountMaximo) && (len(o.SolutionItems) > len(o.OptimaItems))
}

func getIDs(items []models.Item) []string {
	var ids []string
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func calculateTotal(items []models.Item) int {
	var sum int
	for _, item := range items {
		sum = sum + item.Price
	}
	return sum
}

//NewUseCases trigger usescases
func NewUseCases(OptimaItems []models.Item, SolutionItems []models.Item, FinalItems []string) *UseCases {
	return &UseCases{OptimaItems, SolutionItems, FinalItems}
}
