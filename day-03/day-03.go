package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
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
}
