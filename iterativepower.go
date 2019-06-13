package piscine

import (
	"math"
)

func IterativePower(nb int, power int) int {

	if power < 0 {
		return 0
	}
	x, y := float64(nb), float64(power)
	res := math.Pow(x, y)

	return int(res)

}
