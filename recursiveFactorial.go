package piscine

import (
	"math"
)

func RecursiveFactorial(x int) int {
	if 0 > x {
		return 0
	} else if 0 == x || 1 == x {
		return 1
	} else {
		res := 1
		res = x * RecursiveFactorial(x-1)
		if math.MaxInt32 < res {
			return 0
		}
		return res
	}
}
