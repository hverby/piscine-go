package piscine

func Fibonacci(index int) int {
	var w int
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	u := 0
	v := 1

	for i := 2; i <= index; i++ {
		w = u + v
		u = v
		v = w
	}
	return v
}
