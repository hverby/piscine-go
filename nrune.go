package piscine

func NRune(s string, n int) rune {
	res := []rune(s)
	return res[n-1]
}
