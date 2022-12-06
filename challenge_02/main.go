package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func game(player1 string, player2 string) int {
	//change rock with letter
	wins := make(map[string]string)
	wins["X"] = "C"
	wins["Z"] = "B"
	wins["Y"] = "A"

	loses := make(map[string]string)
	loses["X"] = "B"
	loses["Z"] = "A"
	loses["Y"] = "C"
	var points int
	points = 3
	if wins[player2] == player1 {
		points = 6
	}
	if loses[player2] == player1 {
		points = 0
	}
	return points

}

func main() {

	//points := map[string]uint8{"X": 1, "Y": 2,  "Z": 3}

	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	points := 0
	pointsFromSeed := make(map[string]int)
	pointsFromSeed["Y"] = 2
	pointsFromSeed["X"] = 1
	pointsFromSeed["Z"] = 3

	var totalPoints int
	for scanner.Scan() {
		text := scanner.Text()
		player1 := string(text[0])
		player2 := string(text[2])
		points = game(player1, player2)
		totalPoints += points
		totalPoints += pointsFromSeed[player2]
		//If empty line
		if err != nil {
			panic("ahi")
		}
	}
	fmt.Println(totalPoints)
}
