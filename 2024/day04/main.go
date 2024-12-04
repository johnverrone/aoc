package main

import (
	"fmt"
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

func main() {
	in := util.ParseInput("")
	in = strings.TrimSpace(in)

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
			rStart := ""
			lStart := ""
			if r > c {
				rStart = strconv.Itoa(r-c) + ",0"
			} else {
				rStart = "0," + strconv.Itoa(c-r)
			}
			if r+c <= len(rows[0])-1 {
				lStart = "0," + strconv.Itoa(r+c)
			} else if c == len(rows[0])-1 {
				lStart = strconv.Itoa(r) + "," + strconv.Itoa(c)
			} else {
				lStart = strconv.Itoa(r-(len(rows[0])-1-c)) + "," + strconv.Itoa(len(rows[0])-1)
			}
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

	search := strings.Join(words, " ")
	count := strings.Count(search, "XMAS")
	count += strings.Count(search, "SAMX")
	fmt.Println("Part 1:", count)

	// part 2
	// fmt.Println("Part 2:", sum)
}
