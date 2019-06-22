package piscine


func DivMod(a int, b int, div *int, mod *int){
	*div= int(a/b)
	*mod= a%b
}
