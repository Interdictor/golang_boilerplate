package fizzbuzz

import "fmt"

type Rule struct {
	Number int
	Word   string
}

var rules = []Rule{
	{3, "Fizz"},
	{5, "Buzz"},
	{7, "Bazz"},
}

func FizzBuzz(number int) string {
	result := ""

	for _, v := range rules {
		if number%v.Number == 0 {
			result += v.Word
		}
	}

	if result == "" {
		return fmt.Sprint(number)
	}

	return result
}
