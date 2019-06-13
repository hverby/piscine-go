package piscine

func IterativeFactorial(x int) int {
	if x == 0 {
		return 0
	}
	if x > 1 {
		return x * IterativeFactorial(x-1)
	}
	if x == 1 {
		return x * x
	}
	return 0
}
