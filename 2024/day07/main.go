package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

type Equation struct {
	val    int
	inputs []int
}

func main() {
	in := util.ParseInput("")
	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")

	equations := setup(lines)

	// part 1
	sum := 0
	for _, eq := range equations {
		oc := len(eq.inputs) - 1
		ops := buildOperatorString(oc, []rune{'*', '+'})

		for _, o := range ops {
			ans := solveEq(eq.inputs, o)
			if ans == eq.val {
				sum += eq.val
				break
			}
		}
	}
	fmt.Printf("Part 1: %d\n", sum)

	// part 2
	sum = 0
	for _, eq := range equations {
		oc := len(eq.inputs) - 1
		ops := buildOperatorString(oc, []rune{'|', '*', '+'})

		for _, o := range ops {
			ans := solveEq(eq.inputs, o)
			if ans == eq.val {
				sum += eq.val
				break
			}
		}
	}
	fmt.Printf("Part 2: %d\n", sum)
}

func setup(lines []string) []Equation {
	var equations []Equation
	for _, line := range lines {
		parts := strings.Split(line, ":")
		val := util.MustInt(strings.TrimSpace(parts[0]))
		equation := []int{}
		equationString := strings.TrimSpace(parts[1])
		equationStrings := strings.Fields(equationString)
		for _, es := range equationStrings {
			equation = append(equation, util.MustInt(es))
		}

		equations = append(equations, Equation{
			val:    val,
			inputs: equation,
		})
	}
	return equations
}

func solveEq(nums []int, ops []rune) int {
	c := nums[:]
	var ans int
	for op := 0; len(c) >= 2; op++ {
		if ops[op] == '|' {
			ans = util.MustInt(strconv.Itoa(c[0]) + strconv.Itoa(c[1]))
		} else if ops[op] == '+' {
			ans = c[0] + c[1]
		} else {
			ans = c[0] * c[1]
		}
		c = append([]int{ans}, c[2:]...)
	}
	return ans
}

func buildOperatorString(count int, chars []rune) [][]rune {
	result := [][]rune{}
	if count == 0 {
		return result
	}

	var generate func([]rune, int)
	generate = func(perm []rune, index int) {
		if index == count {
			result = append(result, append([]rune{}, perm...))
			return
		}

		for _, char := range chars {
			generate(append(perm, char), index+1)
		}
	}

	generate([]rune{}, 0)
	return result
}
