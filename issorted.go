package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	res := 1
	for k,v := range tab{
		if k != len(tab) -1 {
			if f(v, tab[k+1]) == -1{
				res++
			}
		}
	}
	return res == len(tab)
}