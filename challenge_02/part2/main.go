package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Y draw, X lose, Z win
func game(player1 string, player2 string) int {
	//change rock with letter
	wins := make(map[string]string)
	wins["B"] = "A"
	wins["C"] = "B"
	wins["A"] = "C"

	loses := make(map[string]string)
	loses["A"] = "B"
	loses["B"] = "C"
	loses["C"] = "A"

	var points int
	pointsFromGame := 3
	pointsFromSeed := make(map[string]int)
	pointsFromSeed["A"] = 1
	pointsFromSeed["B"] = 2
	pointsFromSeed["C"] = 3

	points = pointsFromGame + pointsFromSeed[player1]

	if player2 == "X" {
		pointsFromGame := 0
		points = pointsFromGame + pointsFromSeed[wins[player1]]
	}
	if player2 == "Z" {
		pointsFromGame := 6
		points = pointsFromGame + pointsFromSeed[loses[player1]]
	}

	if loses[player2] == player1 {
		points = 0
	}
	return points

}

func main() {

	//points := map[string]uint8{"X": 1, "Y": 2,  "Z": 3}

	f, err := os.Open("../data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var totalPoints int
	for scanner.Scan() {
		text := scanner.Text()
		player1 := string(text[0])
		player2 := string(text[2])
		points := game(player1, player2)
		totalPoints += points
		//If empty line
		if err != nil {
			panic("ahi")
		}
	}
	fmt.Println(totalPoints)
}
