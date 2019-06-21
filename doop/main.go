package main

import (
	"fmt"
	"os"
	"strconv"
)

func isSigne(str string, arr[]string) bool{
	for _,v := range arr{
		if str == v{
			return true
		}
	}
	return false
}

func main(){
	signe := []string{"+","*","-","/","%"}
	args := os.Args[1:]
	if len(args) != 3 {

	}
	if isSigne(args[1], signe){
		nb1, err := strconv.ParseInt(args[0], 10, 64)
		nb2, err2 := strconv.ParseInt(args[0], 10, 64)
		if err == nil && err2 == nil{
			switch args[1] {
			case "+":
				fmt.Println(int(nb1) + int(nb2))
			case "-":
				fmt.Println(int(nb1) - int(nb2))
			case "/":
				if nb2 == 0 {
					fmt.Println("No division by 0")
				}else{
					fmt.Println(int(nb1) / int(nb2))
				}
			case "%":
				if nb2 == 0 {
					fmt.Println("No modulo by 0")
				}else{
					fmt.Println(int(nb1) % int(nb2))
				}

			}
		}

	}else{
		fmt.Println(0)
	}
}