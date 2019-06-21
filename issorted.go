package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	res := 1
	res2 := 1
	res3 := 1
	for k,v := range tab{
		if k != len(tab) -1 {
			if f(v, tab[k+1]) < 0{
				res++
			}
			if f(v, tab[k+1]) > 0{
				res2++
			}
			if f(v, tab[k+1]) == 0{
				res3++
			}
		}
	}
	return res == len(tab) || res2 == len(tab) || res3 == len(tab)
}