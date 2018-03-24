package knapsack

// Item is an item we can place in the Knapsack
type Item struct {
	Weight int
	Value  int
}

// Solution is a max value and mapping of item indexes to a count
// of how many of these items to use.
type Solution struct {
	Value int
	Items map[int]int
}

// Bounded finds a Solution of max value given maxWeight, assuming
// we can use each Item only once.
func Bounded(items []Item, maxWeight int) Solution {

	// solutions is a two dimensional array, mapping the max item index
	// to the max weight allowed to the Solution corresponding
	// to that combination of items and weight
	solutions := make([][]Solution, len(items))
	for i := 0; i < len(items); i++ {
		solutions[i] = make([]Solution, maxWeight+1)
	}

	for i, item := range items {
		for weight := 0; weight <= maxWeight; weight++ {
			var lastSolution, withoutSolution, currentSolution Solution

			// What was the solution not including this item?
			if i > 0 {
				lastSolution = solutions[i-1][weight]
			}

			// What is the solution without the weight the current item
			// provides?
			if i > 0 && (weight-item.Weight) >= 0 {
				withoutSolution = solutions[i-1][weight-item.Weight]
			}

			potentialValue := item.Value + withoutSolution.Value

			if item.Weight <= weight && potentialValue > lastSolution.Value {
				// If we *can* include this item, and it gives us more
				// value, then include it

				// Copy the withoutSolution
				currentSolution.Items = map[int]int{}
				for k, v := range withoutSolution.Items {
					currentSolution.Items[k] = v
				}

				// Add ourselves, set the value, assign to matrix
				currentSolution.Items[i] = 1
				currentSolution.Value = potentialValue
				solutions[i][weight] = currentSolution

			} else {
				// Otherwise, use the solution that didn't consider this
				// item.
				solutions[i][weight] = lastSolution

			}
		}
	}

	return solutions[len(items)-1][maxWeight]
}
