package puzzles

func bgHelper(orig string, word string, letters string, seen map[string]bool) string {
	var (
		i                    int
		letter               rune
		minWord, thisMinWord string
	)

	key := word + "|" + letters

	if seen[key] {
		return ""
	}

	seen[key] = true

	// If the word we have so far is longer than the original
	// and it's lexigraphically before it, there is no solution
	// that we're going to find in this path. Stop early.
	if len(word) >= len(orig) && word <= orig {
		return ""
	}

	// If there are no more letters, there is nothing else to do.
	// Return this word.
	if len(letters) < 1 {
		return word
	}

	// Choose one letter to add to the word for each recursive call
	for i, letter = range letters {

		// Get the min word with this letter added
		thisMinWord = bgHelper(
			orig, word+string(letter), letters[:i]+letters[i+1:], seen,
		)

		// If it's empty, ignore
		if thisMinWord == "" {
			continue
		}

		if minWord == "" || thisMinWord < minWord {
			minWord = thisMinWord
		}
	}

	return minWord
}

// BiggerIsGreater returns the next lexigraphically sorted word we can
// get by swapping one or more of the letters of s.
func BiggerIsGreater(s string) string {
	var (
		nextWord string
		i        int
		letters  string
		word     string
		ok       bool
	)

	seen := map[string]bool{}

	// Short cut. If every character of the string is lexigraphically less or
	// equal than the prior character in the string, there is no answer.
	for i = 1; i < len(s); i++ {
		if s[i-1] < s[i] {
			ok = true
			break
		}
	}

	if !ok {
		return "no answer"
	}

	for i = 1; i <= len(s); i++ {
		word = s[:len(s)-i]
		letters = s[len(s)-i:]

		nextWord = bgHelper(s, word, letters, seen)
		if nextWord != "" {
			return nextWord
		}
	}

	return "no answer"
}
