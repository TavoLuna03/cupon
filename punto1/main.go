package main

type Item struct {
	ID    string
	Price int
}

type OptimaItems struct {
	OptimaItems   []Item
	SolutionItems []Item
	FinalItems    []string
}

func (o *OptimaItems) calculate(items []Item, amount int) []string {
	o.calculateRec(items, amount, false)
	return o.FinalItems
}

func (o *OptimaItems) calculateRec(items []Item, ammountMaximo int, optima bool) {
	//validar si es solucion optima
	if optima {
		// comprobar si es mejor solucion que la ultima
		if o.validateOptimalSolution(ammountMaximo) {
			// actualizar optimaItems con solucion
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

func (o *OptimaItems) validateExistense(itemValidate Item) bool {
	for _, item := range o.SolutionItems {
		if item.ID == itemValidate.ID {
			return true
		}
	}
	return false
}

func (o *OptimaItems) validateOptimalSolution(ammountMaximo int) bool {
	totalSolution := calculateTotal(o.SolutionItems)
	totalOptima := calculateTotal(o.OptimaItems)

	return (totalSolution > totalOptima) && (totalSolution < ammountMaximo) && (len(o.SolutionItems) > len(o.OptimaItems))
}
func getIDs(items []Item) []string {
	var ids []string
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func calculateTotal(items []Item) int {
	var sum int
	for _, item := range items {
		sum = sum + item.Price
	}
	return sum
}
