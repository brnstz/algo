package puzzles

import "fmt"

var (
	words = map[int]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
	}

	specialMinuteAfter = map[int]string{
		0: "o' clock",
	}

	specialMinuteBefore = map[int]string{
		15: "quarter past",
		30: "half past",
		45: "quarter to",
	}
)

func minuteToWord(minute int) string {
	var (
		word, unit, pastTo string
	)

	if minute == 1 {
		unit = "minute"
	} else {
		unit = "minutes"
	}

	if minute < 30 {
		pastTo = "past"
	} else {
		minute = 60 - minute
		pastTo = "to"
	}

	if minute <= 20 {
		word = words[minute]
	} else {
		word = words[minute-minute%10] + " " + words[minute%10]
	}

	return fmt.Sprintf("%v %v %v", word, unit, pastTo)
}

func hourToWord(hour, minute int) string {
	if minute > 30 {
		if hour < 12 {
			hour++
		} else {
			hour = 1
		}
	}

	return words[hour]
}

// TimeInWords returns the time in a string
func TimeInWords(hour, minute int) string {
	var (
		minuteWord, hourWord string
		exists               bool
	)

	hourWord = hourToWord(hour, minute)

	minuteWord, exists = specialMinuteAfter[minute]
	if exists {
		return fmt.Sprintf("%v %v", hourWord, minuteWord)
	}

	minuteWord, exists = specialMinuteBefore[minute]
	if exists {
		return fmt.Sprintf("%v %v", minuteWord, hourWord)
	}

	minuteWord = minuteToWord(minute)

	return fmt.Sprintf("%v %v", minuteWord, hourWord)
}
