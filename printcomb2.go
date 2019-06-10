package piscine

import "fmt"
import "strconv"

func PrintComb2()  {
	for i:=0;i<10;i++{
		for j:=0;j<10;j++{
			for k:=0;k<10;k++{
				for l:=0;l<10;l++{
					nb1:= strconv.Itoa(i) + strconv.Itoa(j)
					nb2:= strconv.Itoa(k) + strconv.Itoa(l)
					if(nb1 < nb2){
						fmt.Print(nb1 + " " +nb2)
						if nb1 < "98" {
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
