package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var calories int
	var maxCalories1 = 0
	var maxCalories2 = 0
	var maxCalories3 = 0

	for scanner.Scan() {
		text := scanner.Text()

		snack, err := strconv.Atoi(text)
		//If empty line
		if err != nil {
			if calories > maxCalories3 {
				maxCalories3 = calories
			}
			if maxCalories3 > maxCalories2 {
				maxCalories3, maxCalories2 = maxCalories2, maxCalories3
			}
			if maxCalories2 > maxCalories1 {
				maxCalories2, maxCalories1 = maxCalories1, maxCalories2
			}
			calories = 0
		}
		calories += snack
	}

	fmt.Println(maxCalories1 + maxCalories2 + maxCalories3)

}
