package computor

import (
	"fmt"
	"math"
	"regexp"
)

// Compute resolve 2nd degree operations
func Compute(text string) {
	fmt.Printf(text)
}

func abs(number float64) float64 {
	if number < 0 {
		return -number
	}
	return number
}

func sqrt(number float64) float64 {
	if number == 0 || math.IsNaN(number) || math.IsInf(number, 1) {
		return number
	}

	x := number / 2.0
	val := 0.0

	for true {
		val = (x + (number / x)) / 2
		if abs(val-x) < .0000001 {
			return val
		}
		x = val
	}
	return x
}

func discriminant(a, b, c float64) float64 {
	return b*b - 4*a*c
}

func checkRules(equation string) bool {
	match, err := regexp.Match("", []byte(equation))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Matched ? : ", match)
	return true
}
