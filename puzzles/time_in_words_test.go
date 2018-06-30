package puzzles

import (
	"fmt"
	"testing"
)

type twordTest struct {
	hour, minute int
	expected     string
}

func TestTimeInWords(t *testing.T) {
	x := []twordTest{
		{5, 0, "five o' clock"},
		{6, 30, "half past six"},
		{6, 31, "twenty nine minutes to seven"},
		{5, 1, "one minute past five"},
		{5, 10, "ten minutes past five"},
		{5, 40, "twenty minutes to six"},
		{5, 45, "quarter to six"},
		{5, 47, "thirteen minutes to six"},
		{5, 28, "twenty eight minutes past five"},
		{12, 47, "thirteen minutes to one"},
		{12, 00, "twelve o' clock"},
	}

	for _, twt := range x {
		actual := TimeInWords(twt.hour, twt.minute)
		fmt.Println(actual)
		if actual != twt.expected {
			t.Fatalf("expected %v but got %v for %v %v", twt.expected,
				actual, twt.hour, twt.minute,
			)
		}
	}
}
