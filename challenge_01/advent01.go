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
	var maxCalories = 0
	for scanner.Scan() {
		text := scanner.Text()

		snack, err := strconv.Atoi(text)
		//If error is different from nil then I found an empty line
		if err != nil {
			if calories > maxCalories {
				maxCalories = calories
			}
			calories = 0
		}
		fmt.Println(snack)
		calories += snack
	}

	fmt.Println(maxCalories)

}
