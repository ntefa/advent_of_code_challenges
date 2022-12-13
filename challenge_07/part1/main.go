package main

import (
	"advent06_1/utils"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var s string

func main() {
	lines := strings.Split(s, "\n")
	//root := utils.MakeDir("root", nil)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		s1 := strings.Split(line, " ")
		//if command
		if s1[0] == "$" {
			command := s[1]
			if s1[1] == "ls" {
				j := i + 1
				for j < len(lines) && strings.Split(lines[j], "\n")[0] != "$" {
					s2 := strings.Split(lines[j], "\n")
					part1, part2 := s2[1], s2[2]
					//if directory
					if part1 == "dir" {
						fmt.Println("")
					} else {
						size, _ := strconv.Atoi(part1)
						file := utils.File{Name: part1, Size: size}
						//link to directory somehow
					}
				}
			}
		}
	}

}
