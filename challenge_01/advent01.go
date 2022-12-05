package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var s []string
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	var currentElement string
	previousElement := ""
	for i := 0; i < len(s); i++ {
		if i != 0 {
			previousElement = s[i-1]
		}
		currentElement = s[i]
		if currentElement == " " && previousElement == " " {
			fmt.Println(i)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
