package algo

import (
	"testing"
)

func TestNonDivisibleSubset(t *testing.T) {
	x := MaxNonDivisibleSubset(3, []int{1, 7, 2, 4})
	if x != 3 {
		t.Fatalf("expected %v but got %v", 3, x)
	}
}

func TestNonDivisibleSubset1(t *testing.T) {
	x := MaxNonDivisibleSubset(11, []int{582740017, 954896345, 590538156, 298333230, 859747706, 155195851, 331503493, 799588305, 164222042, 563356693, 80522822, 432354938, 652248359})
	if x != 11 {
		t.Fatalf("expected %v but got %v", 11, x)
	}
}

func TestNonDivisibleSubset2(t *testing.T) {
	x := MaxNonDivisibleSubset(1, []int{1, 2, 3, 4, 5})
	if x != 1 {
		t.Fatalf("expected %v but got %v", 1, x)
	}
}

func TestNonDivisibleSubset3(t *testing.T) {
	x := MaxNonDivisibleSubsetIterative(3, []int{1, 7, 2, 4})
	if x != 3 {
		t.Fatalf("expected %v but got %v", 3, x)
	}
}

func TestNonDivisibleSubset4(t *testing.T) {
	x := MaxNonDivisibleSubsetIterative(5, []int{770528134, 663501748, 384261537, 800309024, 103668401, 538539662, 385488901, 101262949, 557792122, 46058493})
	if x != 6 {
		t.Fatalf("expected %v but got %v", 6, x)
	}
}

func TestNonDivisibleSubset5(t *testing.T) {
	x := MaxNonDivisibleSubsetIterative(11, []int{582740017, 954896345, 590538156, 298333230, 859747706, 155195851, 331503493, 799588305, 164222042, 563356693, 80522822, 432354938, 652248359})
	if x != 11 {
		t.Fatalf("expected %v but got %v", 11, x)
	}
}