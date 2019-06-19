package main

import (
	"./z01"
	"os"
)

func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven() int {
	lengthOfArg := os.Args[1:]
	if len(lengthOfArg) % 2 == 0{
		return 1
	}
	return 0
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	
	if isEven() == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}