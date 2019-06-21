package piscine

func Map(f func(int) bool, arr []int) [6]bool {

	var tab [6]bool
	for k,v := range arr{
		if v >= 0 {
			tab[k] = f(v)

		}
	}
	return tab
}