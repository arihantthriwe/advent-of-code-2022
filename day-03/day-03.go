package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimSpace(input)

	firstPart := make(map[string]int)

	var error int

	for _, row := range strings.Split(input, "\n") {
		for i, line := range row {
			s := string(line)
			n, _ := strconv.Atoi(s)
			if i < len(row)/2 {
				firstPart[s]++
			} else if firstPart[s] > 0 {
				if n >= 97 && n <= 122 {
					error += n - 97 + 1
				} else {
					error += n - 65 + 27
				}
			}
		}
	}
	fmt.Println(error)
}
