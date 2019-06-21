package piscine

func CountIf(f func(string) bool, tab []string) int {
	nb := 0
	for _,v := range tab{
		if f(v){
			nb++
		}
	}
	return nb
}