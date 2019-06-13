package piscine

import "math"

func Sqrt(nb int) int {
	if nb%2 == 0 {
		x := math.Sqrt(float64(nb))
		return int(x)
	}

	return 0

}
