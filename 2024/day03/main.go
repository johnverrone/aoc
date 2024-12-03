package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5)) `

func main() {
	in := util.ParseInput("")

	// part 1
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	sum := 0
	multStrings := re.FindAllString(in, -1)
	for _, s := range multStrings {
		s = s[4 : len(s)-1]
		nums := strings.Split(s, ",")
		a, b := util.MustInt(nums[0]), util.MustInt(nums[1])

		sum += a * b
	}
	fmt.Println("Part 1:", sum)

	// part 2
	re = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	sum = 0
	multStrings = re.FindAllString(in, -1)
	enabled := true
	for _, s := range multStrings {
		if s == "do()" {
			enabled = true
		}
		if s == "don't()" {
			enabled = false
		}
		if enabled && s[0] == 'm' {
			s = s[4 : len(s)-1]
			nums := strings.Split(s, ",")
			a, b := util.MustInt(nums[0]), util.MustInt(nums[1])
			sum += a * b
		}
	}
	fmt.Println("Part 2:", sum)
}
