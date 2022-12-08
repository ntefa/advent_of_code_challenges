package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func findBadge(ruckSackGroup []string) string {
	charCounter := make(map[rune]int)
	for _, rucksack := range ruckSackGroup {
		hash := make(map[rune]bool)
		for _, e := range rucksack {
			if !hash[e] {
				hash[e] = true
				charCounter[e]++
			}
			if charCounter[e] == len(ruckSackGroup) {
				return string(e)
			}
		}
	}
	return ""
}

func calculatePriority() map[string]int {
	priority := make(map[string]int)
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	priority[""] = 0
	for i := 0; i < 26; i++ {
		priority[string(lowercase[i])] = i + 1
		priority[string(uppercase[i])] = i + 27
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
	counter := 1 //counter sets the groups of three
	var ruckSackGroup []string
	for scanner.Scan() {
		rucksack := scanner.Text()
		ruckSackGroup = append(ruckSackGroup, rucksack)
		//If err
		if counter == 3 {
			badge := findBadge(ruckSackGroup)
			fmt.Println(ruckSackGroup)
			fmt.Println(badge)
			//fmt.Println(badge)
			totalPoints += priority[badge]
			counter = 0
			ruckSackGroup = nil

		}

		counter++
	}
	fmt.Println(totalPoints)
}
