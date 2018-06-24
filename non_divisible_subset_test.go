package algo

import (
	"testing"
)

func TestNonDivisibleSubset(t *testing.T) {
	x := MaxNonDivisibleSubset(3, []int{1, 7, 2, 4})
	t.Fatal(x)
}

func TestNonDivisibleSubset1(t *testing.T) {
	x := MaxNonDivisibleSubset(11, []int{582740017, 954896345, 590538156, 298333230, 859747706, 155195851, 331503493, 799588305, 164222042, 563356693, 80522822, 432354938, 652248359})
	t.Fatal(x)
}
