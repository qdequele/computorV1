package computor

import "math"

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

func Discriminant(a, b, c float64) float64 {
	return b*b - 4*a*c
}
