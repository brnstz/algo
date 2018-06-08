package coins

import (
	"fmt"
)

// changeR is a helper function for ChangeRecursive
func changeR(coins []int, amount int, cache map[string]int) int {
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
	combos = changeR(coins[:len(coins)-1], amount, cache) +
		changeR(coins, amount-coins[len(coins)-1], cache)

	cache[key] = combos

	return combos
}

// ChangeRecursive returns the number of possible coin combinations, using a
// cache of previously solved sub-problems
func ChangeRecursive(coins []int, amount int) int {
	cache := map[string]int{}
	return changeR(coins, amount, cache)
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
			combos[amount] += combos[amount-coin]
		}

	}

	return combos[totalAmount]
}

// ChangeLimited returns all possible permutations of change for amount
// with a limited number of coins
func ChangeLimited(coins []int, used []int, amount int) [][]int {
	var (
		combos   [][]int
		newCoins []int
		newUsed  []int
	)

	// We've found a way to make change
	if amount == 0 {
		return [][]int{used}
	}

	// If amount is negative, we've gone too far
	if amount < 0 {
		return nil
	}

	// Iterate with one coin chosen
	for i := range coins {
		newCoins = make([]int, len(coins))
		newUsed = make([]int, len(used))

		copy(newCoins, coins)
		copy(newUsed, used)

		newCoins = append(newCoins[0:i], newCoins[i+1:]...)
		newUsed = append(newUsed, coins[i])

		combos = append(
			combos,
			ChangeLimited(newCoins, newUsed, amount-coins[i])...,
		)
	}

	return combos

}
