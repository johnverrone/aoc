package main

import (
	"fmt"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

type Dir int

const (
	UP Dir = iota
	RIGHT
	DOWN
	LEFT
)

type point struct {
	r, c int
}

type Guard struct {
	r, c    int
	dir     Dir
	m       []string
	visited map[point]int
}

func NewGuard(r, c int, m []string) *Guard {
	return &Guard{
		r, c, UP, m, map[point]int{{r, c}: 1},
	}
}

func (g *Guard) move() {
	for {
		for g.canMove() {
			// move
			switch g.dir {
			case UP:
				g.r--
			case LEFT:
				g.c--
			case RIGHT:
				g.c++
			case DOWN:
				g.r++
			}
			g.visited[point{g.r, g.c}] = 1
		}
		// check bounds & turn
		if (g.dir == UP && g.r == 0) || (g.dir == DOWN && g.r == len(g.m)-1) || (g.dir == RIGHT && g.c == len(g.m[0])-1) || (g.dir == LEFT && g.c == 0) {
			break
		}
		if g.dir >= 3 {
			g.dir = 0
		} else {
			g.dir++
		}
	}
}

func (g *Guard) canMove() bool {
	switch g.dir {
	case UP:
		return g.r >= 0 && g.m[g.r-1][g.c] != '#'
	case LEFT:
		return g.c-1 >= 0 && g.m[g.r][g.c-1] != '#'
	case RIGHT:
		return g.c+1 <= len(g.m[0])-1 && g.m[g.r][g.c+1] != '#'
	case DOWN:
		return g.r+1 <= len(g.m)-1 && g.m[g.r+1][g.c] != '#'
	}
	panic("unhandled move")
}

func main() {
	in := util.ParseInput("")
	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")
	// visited := [][]bool{}

	// part 1
	var g *Guard
	for r, row := range lines {
		idx := strings.Index(row, "^")
		if idx >= 0 {
			fmt.Printf("starting loc %d, %d\n", r, idx)
			g = NewGuard(r, idx, lines)
		}
	}

	g.move()

	fmt.Printf("Part 1: %d\n", len(g.visited))

	// part 2
	p2 := 0
	fmt.Println("Part 2:", p2)
}
