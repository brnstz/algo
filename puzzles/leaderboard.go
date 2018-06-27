package puzzles

// https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem

// Leaderboard returns the place among topScores of each value in
// userScores.
func Leaderboard(topScores []int, userScores []int) []int {
	var (
		worstPlace, bestPlace, place   int
		userScore, topScore, lastScore int
		places                         []int
		found                          bool
	)

	// Set the worst possible place as max int
	worstPlace = int(^uint(0) >> 1)
	bestPlace = worstPlace

	for _, userScore = range userScores {

		// Reset the place setting every time we check a score
		place = 1
		found = false

		// Keep track of last score so we know when to increment place
		lastScore = int(^uint(0) >> 1)

		for _, topScore = range topScores {

			// If our score is better than topScore, then we can
			// record our place.
			if userScore >= topScore {

				// If we get a lower score later on, we only want to record our
				// best score
				if place < bestPlace {
					bestPlace = place
				}

				// Append our best score to places
				places = append(places, bestPlace)

				// We found a place, so won't need to append after we
				// we break.
				found = true
				break
			}

			// If last score is different, then increment the place value.
			if topScore != lastScore {
				place++
			}

			// Keep track of lastScore
			lastScore = topScore
		}

		// If we didn't find a place, we want to use either our best possible
		// place or the current place.
		if !found {

			if bestPlace == worstPlace {
				places = append(places, place)
			} else {
				places = append(places, bestPlace)
			}
		}

	}

	return places
}
