package piscine

import "fmt"
import "strconv"

func PrintComb2()  {
	for x:=0;x<10;x++{
		for y:=0;y<10;y++{
			for z:=0;z<10;z++{
				for l:=0;l<10;l++{
					nbr1:= strconv.Itoa(x) + strconv.Itoa(y)
					nbr2:= strconv.Itoa(z) + strconv.Itoa(l)
					if(nbr1 < nbr2){
						fmt.Print(nbr1 + " " +nbr2)
						
						
						if nbr1 < "98" {
							fmt.Print(", ")
						}else {
							fmt.Print("\n")
						}
					}
				}
			}
		}
	}
}
