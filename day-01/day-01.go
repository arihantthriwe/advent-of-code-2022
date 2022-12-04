package day01

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Day01() {
	input = strings.TrimSpace(input)

	var max int
	var calories []int

	for _, elf := range strings.Split(input, "\n\n") {
		var sum int
		for _, row := range strings.Split(elf, "\n") {
			num, _ := strconv.Atoi(row)
			sum += num
		}
		if sum > max {
			max = sum
		}
		calories = append(calories, sum)
	}

	fmt.Println(max)
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	fmt.Println(calories[0] + calories[1] + calories[2])
}
