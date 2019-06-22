package main

import (
	"fmt"
	"os"
)

func cat(){
	args := os.Args[1:]
	if len(args) == 0 {
		 fmt.Println("Hello")
		 fmt.Println("Hello")
	}
	if len(args) > 0 {
		for i := 0; i < len(args) ;i++{
			file, err := os.Open(args[i])
			if err != nil{
				fmt.Printf("the mistake is %v\n", err.Error())
			}
			arr := make([]byte, 500)
			file.Read(arr)
			fmt.Println(string(arr))
		}
	}
}


func main()  {
	cat()
}