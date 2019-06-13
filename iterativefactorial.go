package piscine

/*import (
	"fmt"
)*/
func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 0
	}
	if nb > 1 {
		return nb * IterativeFactorial(nb-1)
	}
	if nb == 1 {
		return nb * nb
	}
	return 0
}

/*func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}*/
