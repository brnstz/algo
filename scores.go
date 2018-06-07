package algo

import "log"

var (
	// AllPossiblePoints FIXME
	AllPossiblePoints = []int{7, 6, 3, 2}
	pointCache        = map[int]int{}
)

// PointCombinations returns the number of ways we can arrive an points
// given possiblePoints
func PointCombinations(points int, possiblePoints []int) int {
	log.Printf("points: %v\n", points)
	var (
		combinations int
		exists       bool
		pointsNow    int
	)

	// Check cache first
	combinations, exists = pointCache[points]
	if exists {
		return combinations
	}

	for _, poss := range possiblePoints {
		pointsNow = points
		for pointsNow-poss >= 0 {
			pointsNow = pointsNow - poss

			if pointsNow == 0 {
				log.Printf("%v %v ++\n", pointsNow, poss)
				combinations++
				continue
			}

			if len(possiblePoints) > 1 {
				combinations += PointCombinations(
					pointsNow,
					possiblePoints[:len(possiblePoints)-1],
				)
			}
		}
	}
	log.Println()

	pointCache[points] = combinations

	return combinations
}
