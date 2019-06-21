package piscine

func Map(f func(int) bool, arr []int) []bool {

	tab := make([]bool, len(arr))

	for k,v := range arr{
		if v >= 0 {
			tab[k] = f(v)
		}
	}
	return tab
}