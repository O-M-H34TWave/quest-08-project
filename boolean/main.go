package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg := len(os.Args[1:])
	evenMsg := "I have an even number of arguments"
	oddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) {
		printStr(evenMsg)
	} else {
		printStr(oddMsg)
	}
}
