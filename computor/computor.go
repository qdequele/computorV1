package computor

import (
	"fmt"
	"regexp"
)

// Compute resolve 2nd degree operations
func Compute(text string) {
	fmt.Printf(text)
}

func checkRules(equation string) bool {
	match, err := regexp.Match("", []byte(equation))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Matched ? : ", match)
	return true
}
