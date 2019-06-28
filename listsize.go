package piscine 

func ListSize(l *List) int {
	it := l.Head
	i := 0
	for it != nil {
		i++
		it = it.Next
	}
	return i
}