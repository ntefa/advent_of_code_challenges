package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var s string

type Stacks [][]string

func Split(r rune) bool {
	return r == 'm' || r == 'f' || r == 't'
}

func (stacks *Stacks) initData() [][]string {
	stack := []string{"Z", "J", "G"}
	*stacks = append(*stacks, stack)
	stack = []string{"Q", "L", "R", "P", "W", "F", "V", "C"}
	*stacks = append(*stacks, stack)
	stack = []string{"F", "P", "M", "C", "L", "G", "R"}
	*stacks = append(*stacks, stack)
	stack = []string{"L", "F", "B", "W", "P", "H", "M"}
	*stacks = append(*stacks, stack)
	stack = []string{"G", "C", "F", "S", "V", "Q"}
	*stacks = append(*stacks, stack)
	stack = []string{"W", "H", "J", "Z", "M", "Q", "T", "L"}
	*stacks = append(*stacks, stack)
	stack = []string{"H", "F", "S", "B", "V"}
	*stacks = append(*stacks, stack)
	stack = []string{"F", "J", "Z", "S"}
	*stacks = append(*stacks, stack)
	stack = []string{"M", "C", "D", "P", "F", "H", "B", "T"}
	*stacks = append(*stacks, stack)
	return *stacks
}

func (stacks *Stacks) pop(from int) string {

	stack := (*stacks)[from]
	f := len(stack)
	lastEle := (stack)[f-1]
	stack = stack[:f-1]
	(*stacks)[from] = stack
	return lastEle
}
func (stacks *Stacks) push(to int, element string) {

	stack := (*stacks)[to]
	stack = append(stack, element)
	(*stacks)[to] = stack
}

func (stacks *Stacks) moveStacks(from, to, quantity int) {
	var lastEle string
	for i := 0; i < quantity; i++ {
		lastEle = stacks.pop(from)
		stacks.push(to, lastEle)
	}
}

func parseInput(input string) []string {
	temp := strings.ReplaceAll(input, "move", "")
	temp = strings.ReplaceAll(temp, "from", "")
	temp = strings.ReplaceAll(temp, "to", "")

	res := strings.SplitN(temp, "  ", -1)

	return res
}
func (stacks *Stacks) computeMsg() string {
	var msg string
	for _, stack := range *stacks {
		msg = msg + stack[len(stack)-1]
	}
	return msg
}

func main() {
	var myStacks Stacks
	myStacks.initData()

	moves := strings.Split(s, "\n")
	for _, move := range moves {
		res := parseInput(move)
		quantity, _ := strconv.Atoi(strings.ReplaceAll(res[0], " ", ""))
		from, _ := strconv.Atoi(strings.ReplaceAll(res[1], " ", ""))
		to, _ := strconv.Atoi(strings.ReplaceAll(res[2], " ", ""))

		myStacks.moveStacks(from-1, to-1, quantity)
	}
	fmt.Println(myStacks.computeMsg())
}
