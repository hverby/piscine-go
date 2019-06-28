package piscine 

func ListSize(l *List) int {
	it := l.Head
	i := 0
	for it != nil {
		i++
		fmt.Println(it.Data)
		it = it.Next
	}
	return i
}