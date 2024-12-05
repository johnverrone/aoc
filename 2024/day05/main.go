package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

func main() {
	in := util.ParseInput("")
	in = strings.TrimSpace(in)
	ins := strings.Split(in, "\n\n")
	rulesString := strings.Split(ins[0], "\n")
	rules := map[int][]int{}
	for _, r := range rulesString {
		rs := strings.Split(r, "|")
		f := util.MustInt(rs[0])
		l := util.MustInt(rs[1])
		rules[f] = append(rules[f], l)
	}

	update := ins[1]
	updates := [][]int{}
	updatesString := strings.Split(update, "\n")
	for _, u := range updatesString {
		numsString := strings.Split(u, ",")
		nums := []int{}
		for _, s := range numsString {
			s = strings.TrimSpace(s)
			nums = append(nums, util.MustInt(s))
		}
		updates = append(updates, nums)
	}

	// part 1
	sum := 0
	valid := [][]int{}
	invalid := []struct {
		update    []int
		violatedF int
		violatedL int
	}{}
	for _, u := range updates {
		if v, f, l := isValidUpdate(u, rules); v {
			valid = append(valid, u)
		} else {
			invalid = append(invalid, struct {
				update    []int
				violatedF int
				violatedL int
			}{
				update:    u,
				violatedF: f,
				violatedL: l,
			})
		}
	}
	for _, v := range valid {
		sum += v[len(v)/2]
	}
	fmt.Println("Part 1:", sum)

	// part 2
	p2 := 0
	for _, inv := range invalid {
		fixInvalid(&inv.update, rules)
		p2 += inv.update[len(inv.update)/2]
	}
	fmt.Println("Part 2:", p2)
}

func fixInvalid(arr *[]int, rules map[int][]int) {
	v, f, l := isValidUpdate(*arr, rules)
	if v == true {
		return
	}
	swapInvalid(arr, f, l)
	fixInvalid(arr, rules)
}

func swapInvalid(arr *[]int, a int, b int) {
	aIdx := slices.Index(*arr, a)
	bIdx := slices.Index(*arr, b)
	(*arr)[aIdx], (*arr)[bIdx] = (*arr)[bIdx], (*arr)[aIdx]
}

func isValidUpdate(u []int, rules map[int][]int) (bool, int, int) {
	for i := range u {
		if rs, ok := rules[u[i]]; ok {
			// found a rule
			// look for violation
			for _, r := range rs {
				locInUpdate := slices.Index(u, r)
				if locInUpdate >= 0 && locInUpdate < i {
					// invalid update
					return false, u[i], r
				}
			}
		}
	}

	return true, -1, -1
}
