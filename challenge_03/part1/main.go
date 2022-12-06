package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func isEven(num int) bool {
	return num%2 == 0
}
func splitCompartments(rucksack string) (string, string, error) {
	length := len(rucksack)
	if isEven(length) {
		compartment1 := rucksack[:length/2]
		compartment2 := rucksack[length/2:]
		return compartment1, compartment2, nil
	}
	return "", "", errors.New("rucksack length is odd")
}

func intersection(compartment1, compartment2 string) (inter string) {
	hash := make(map[rune]bool)
	for _, e := range compartment1 {
		hash[e] = true
	}
	for _, e := range compartment2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			return string(e)
		}
	}
	return ""
}

func calculatePriority() map[string]int {
	priority := make(map[string]int)
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	counter := 0
	for _, e := range lowercase {
		counter++
		priority[string(e)] = counter
	}
	for _, e := range uppercase {
		counter++
		priority[string(e)] = counter
	}
	return priority

}
func main() {

	//points := map[string]uint8{"X": 1, "Y": 2,  "Z": 3}

	f, err := os.Open("../data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalPoints := 0
	priority := calculatePriority()
	for scanner.Scan() {
		rucksack := scanner.Text()
		compartment1, compartment2, err := splitCompartments(rucksack)
		//If err
		if err != nil {
			log.Fatal(err)
		}

		commonItem := intersection(compartment1, compartment2)
		totalPoints += priority[commonItem]
	}
	fmt.Println(totalPoints)
}
