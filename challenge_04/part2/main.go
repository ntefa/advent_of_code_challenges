package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var s string

func Split(r rune) bool {
	return r == '-' || r == ','
}

func computePairs(pairs []string) ([]int, []int) {
	pair11, _ := strconv.Atoi(pairs[0])
	pair12, _ := strconv.Atoi(pairs[1])
	pair21, _ := strconv.Atoi(pairs[2])
	pair22, _ := strconv.Atoi(pairs[3])

	pair1 := []int{pair11, pair12}
	pair2 := []int{pair21, pair22}

	return pair1, pair2

}

// func to evaluate wheter pair1 contains pair2
func containsPair(pair1, pair2 []int) bool {
	if pair1[0] <= pair2[0] && pair1[1] >= pair2[1] {
		return true
	}
	return false
}

func doesOverlap(pair1, pair2 []int) bool {
	if pair1[0] > pair2[0] {
		pair1, pair2 = pair2, pair1
	}

	return pair1[1] >= pair2[0]
}

func main() {

	pairs := strings.Split(s, "\n")
	counter := 0
	for _, pair := range pairs {
		assignments := strings.FieldsFunc(pair, Split)
		pair1, pair2 := computePairs(assignments)
		if doesOverlap(pair1, pair2) {
			counter++
			fmt.Println(assignments)
			fmt.Println(counter)
		}

	}
	fmt.Println(counter)

}
