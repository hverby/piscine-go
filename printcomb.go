package piscine

import "fmt"

func PrintComb() {
         
	j:=1
	k:=2
	for i := 0; i <= 7; i++ {
		for j = i + 1; j <= 8; j++ {
			for k = j + 1; k <= 9; k++ {
			       
				fmt.Print(i)
				fmt.Print(j)
				fmt.Print(k)
				if i < 7 {
				fmt.Print(", ")
				}else {
				    fmt.Print("\n")
				}

			}
		}
	}
	
}
