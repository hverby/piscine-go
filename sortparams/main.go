package main

import (
	"os"
	"fmt"
	"sort"
)

func main(){
	args := os.Args[1:]
	
	sort.Slice(args, func(x, y int) bool {
	    return args[x] < args[y]
	})
	
	for _,res:= range args{
		fmt.Println(res)
	}

}
