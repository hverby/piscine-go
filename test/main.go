package main

import "fmt"

func IsNegative(nb int) {
	if nb >= 0{
		fmt.Println("F")
	}else{
		fmt.Println("T")
	}
}

func main()  {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}
