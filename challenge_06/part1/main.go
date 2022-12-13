package main

import (
	"advent06_1/utils"
	_ "embed"
	"fmt"
)

//go:embed data.txt
var s string

func main() {

	for i := 0; i < len(s)-13; i++ {
		currentString := s[i : i+14]
		fmt.Println(currentString)
		if utils.AreStringElementsDifferent(currentString) {
			fmt.Println(i + 14)
			break
		}
	}
}
