package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

var re = regexp.MustCompile("XMAS|SAMX")

func main() {
	in := util.ParseInput(sample)

	// part 1
	words := []string{}
	rows := strings.Split(in, "\n")

	// horizontal words
	words = append(words, rows...)

	// vertical words
	for c := 0; c < len(rows[0]); c++ {
		word := ""
		for r := 0; r < len(rows); r++ {
			word += string(rows[r][c])
		}
		words = append(words, word)
	}

	// diagonal words
	rDiag := map[string]string{}
	lDiag := map[string]string{}
	for r := 0; r < len(rows); r++ {
		for c := 0; c < len(rows[0]); c++ {
			fmt.Printf("trying to find a home for %d, %d\n", r, c)
			rStart := ""
			lStart := ""
			if r > c {
				rStart = strconv.Itoa(r-c) + ",0"
			} else {
				rStart = "0," + strconv.Itoa(c-r)
			}
			if r > len(rows[0])-c {
				lStart = strconv.Itoa(r+c) + ",0"
			} else {
				lStart = "0," + strconv.Itoa(r+c)
			}
			fmt.Printf("\tright goes to %s and left goes to %s\n", rStart, lStart)
			rDiag[rStart] += string(rows[r][c])
			lDiag[lStart] += string(rows[r][c])
		}
	}
	for _, v := range rDiag {
		words = append(words, v)
	}
	for _, v := range lDiag {
		words = append(words, v)
	}

	fmt.Println(len(words))

	search := strings.Join(words, " ")
	matches := re.FindAllStringIndex(search, -1)
	fmt.Println("Part 1:", len(matches))

	// part 2
	// fmt.Println("Part 2:", sum)
}
