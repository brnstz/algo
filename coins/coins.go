package coins

import "fmt"

func _crc(coins []int, amount int, cache map[string]int) int {
	var (
		key    string
		combos int
		exists bool
	)
	key = fmt.Sprintf("%v|%v", len(coins), amount)
	combos, exists = cache[key]
	if exists {
		return combos
	}

	// There is exactly one way to make amount 0
	if amount == 0 {
		return 1
	}

	// There is no way to make a negative amount
	if amount < 0 {
		return 0
	}

	// If there are no coins left but the amount is > 0, there
	// are no solutions
	if len(coins) < 1 && amount > 0 {
		return 0
	}

	// Compute the number of combos both using and not using this coin
	combos = _crc(coins[:len(coins)-1], amount, cache) +
		_crc(coins, amount-coins[len(coins)-1], cache)

	cache[key] = combos

	return combos
}

// ChangeRecursive returns the number of possible coin combinations, using a
// cache of previously solved sub-problems
func ChangeRecursive(coins []int, amount int) int {
	cache := map[string]int{}
	return _crc(coins, amount, cache)
}

// ChangeIterative returns the number of possible coin combinations using an
// iterative bottom up solution
func ChangeIterative(coins []int, totalAmount int) int {
	// We save all possible combos from amount=0 to the actual amount.
	// This will be pre-initialized to 0 for every possibility.
	combos := make([]int, totalAmount+1)

	// There is exactly 1 way to make change for a 0 amount
	combos[0] = 1

	// For every coin
	for _, coin := range coins {

		// For every amount equal to or greater than coin value
		for amount := coin; amount <= totalAmount; amount++ {
			fmt.Printf("before: amount: %v, combos: %v\n", amount, combos)

			combos[amount] += combos[amount-coin]

			fmt.Printf("after: amount: %v, combos: %v\n", amount, combos)

		}

	}

	return combos[totalAmount]
}
