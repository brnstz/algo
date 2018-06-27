package puzzles

import "github.com/brnstz/algo"

// Rect is a rectangle from (X1, Y1) -> (X2, Y2)
type Rect struct {
	X1, Y1, X2, Y2 int
}

// Intersection returns the intersection if r1 and r2 intersect, nil otherwise.
func (r1 *Rect) Intersection(r2 *Rect) *Rect {
	if !(r1.X1 < r2.X2 && r1.X2 > r2.X1 &&
		r1.Y1 < r2.Y2 && r1.Y2 > r2.Y1) {

		return nil
	}

	// Assume there is an intersection and create r3
	r3 := &Rect{
		X1: algo.MaxInt(r1.X1, r2.X1),
		X2: algo.MinInt(r1.X2, r2.X2),

		Y1: algo.MaxInt(r1.Y1, r2.Y1),
		Y2: algo.MinInt(r1.Y2, r2.Y2),
	}

	return r3
}
