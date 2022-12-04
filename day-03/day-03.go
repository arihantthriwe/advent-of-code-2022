package day03

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func Day03() {
	input = strings.TrimSpace(input)

	var error int

	for _, row := range strings.Split(input, "\n") {

		firstPart := make(map[rune]int)

		for i, line := range row {
			if i < len(row)/2 {
				firstPart[line]++
			} else if firstPart[line] > 0 {
				if line >= 'a' && line <= 'z' {
					error += int(line-'a') + 1
				} else {
					error += int(line-'A') + 27
				}
				break
			}
		}
	}
	fmt.Println(error)
	secondPart()
}

func secondPart() {

	chars1 := make(map[rune]int)
	chars2 := make(map[rune]int)

	var sumOfPriorities int

	for i, row := range strings.Split(input, "\n") {

		for _, c := range row {
			if (i+1)%3 == 0 {
				if chars1[c] > 0 && chars2[c] > 0 {
					if c >= 'a' && c <= 'z' {
						sumOfPriorities += int(c-'a') + 1
					} else {
						sumOfPriorities += int(c-'A') + 27
					}
					chars1 = make(map[rune]int)
					chars2 = make(map[rune]int)
					break
				}
			} else if (i+1)%3 == 1 {
				chars1[c]++
			} else if (i+1)%3 == 2 {
				chars2[c]++
			}
		}
	}
	fmt.Println(sumOfPriorities)
}
