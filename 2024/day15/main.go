package main

import (
	"fmt"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<
`

type robot struct {
	x, y int
}

var m [][]byte

func main() {
	in := util.ParseInput("")
	parts := strings.Split(in, "\n\n")
	mstring := strings.Split(parts[0], "\n")
	for i, s := range mstring {
		m = append(m, []byte{})
		m[i] = append(m[i], []byte(s)...)
	}
	moves := parts[1]

	r := &robot{-1, -1}

	for y, line := range m {
		for x, ch := range line {
			if ch == '@' {
				r.x = x
				r.y = y
			}
		}
	}

	fmt.Printf("robot starting at %d %d\n", r.x, r.y)

	for _, ch := range moves {
		switch ch {
		case '<':
			r.moveLeft()
		case '>':
			r.moveRight()
		case '^':
			r.moveUp()
		case 'v':
			r.moveDown()
		}
		// printMap()
	}

	sum := 0
	for y := range m {
		for x := range m[y] {
			if m[y][x] == 'O' {
				sum += 100*y + x
			}
		}
	}
	fmt.Printf("%d\n", sum)
}

func (r *robot) moveLeft() {
	// check for open space left of char
	openSpace := false
	i := r.x - 1
	for ; m[r.y][i] != '#'; i-- {
		if m[r.y][i] == '.' {
			openSpace = true
			break
		}
	}

	if openSpace {
		// everything right of this space should shift
		for j := i; j < r.x; j++ {
			m[r.y][j] = m[r.y][j+1]
		}

		// move player
		r.x--
		m[r.y][r.x+1] = '.'
	}
}

func (r *robot) moveRight() {
	// check for open space right of char
	openSpace := false
	i := r.x + 1
	for ; m[r.y][i] != '#'; i++ {
		if m[r.y][i] == '.' {
			openSpace = true
			break
		}
	}

	if openSpace {
		// everything left of this space should shift
		for j := i; j > r.x; j-- {
			m[r.y][j] = m[r.y][j-1]
		}

		// move player
		r.x++
		m[r.y][r.x-1] = '.'
	}
}

func (r *robot) moveUp() {
	// check for open space above char
	openSpace := false
	i := r.y - 1
	for ; m[i][r.x] != '#'; i-- {
		if m[i][r.x] == '.' {
			openSpace = true
			break
		}
	}

	if openSpace {
		// everything below this space should shift
		for j := i; j < r.y; j++ {
			m[j][r.x] = m[j+1][r.x]
		}

		// move player
		r.y--
		m[r.y+1][r.x] = '.'
	}
}

func (r *robot) moveDown() {
	// check for open space below char
	openSpace := false
	i := r.y + 1
	for ; m[i][r.x] != '#'; i++ {
		if m[i][r.x] == '.' {
			openSpace = true
			break
		}
	}

	if openSpace {
		// everything above this space should shift
		for j := i; j > r.y; j-- {
			m[j][r.x] = m[j-1][r.x]
		}

		// move player
		r.y++
		m[r.y-1][r.x] = '.'
	}
}

func printMap() {
	for y := range m {
		for x := range m[y] {
			fmt.Printf("%c", m[y][x])
		}
		fmt.Println()
	}
}
