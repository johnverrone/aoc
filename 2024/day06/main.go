package main

import (
	"fmt"
	"slices"
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
	r, c      int
	dir       Dir
	m         []string
	visited   map[point][]Dir
	loopCount int
}

func NewGuard(r, c int, m []string) *Guard {
	return &Guard{
		r, c, UP, m, map[point][]Dir{{r, c}: {UP}}, 0,
	}
}

func (g *Guard) move() {
	for {
		for g.canMove() {
			// fmt.Printf("can move %v at: %d, %d\n", g.dir, g.r, g.c)
			//check if we can block
			g.canBlock()
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
			g.visited[point{g.r, g.c}] = append(g.visited[point{g.r, g.c}], g.dir)
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

func (g *Guard) canBlock() {
	yes := false
	switch g.dir {
	case UP:
		for i := 1; g.c+i < len(g.m[0])-1 && g.m[g.r][g.c+i] != '#'; i++ {
			vDir, ok := g.visited[point{g.r, g.c + i}]
			yes = ok && slices.Contains(vDir, RIGHT)
		}
	case LEFT:
		for i := 1; g.r-i >= 0 && g.m[g.r-i][g.c] != '#'; i++ {
			vDir, ok := g.visited[point{g.r - i, g.c}]
			yes = ok && slices.Contains(vDir, UP)
		}
	case RIGHT:
		for i := 1; g.r+i < len(g.m)-1 && g.m[g.r+i][g.c] != '#'; i++ {
			vDir, ok := g.visited[point{g.r + 1, g.c}]
			yes = ok && slices.Contains(vDir, DOWN)
		}
	case DOWN:
		for i := 1; g.c-i >= 0 && g.m[g.r][g.c-1] != '#'; i++ {
			vDir, ok := g.visited[point{g.r, g.c - 1}]
			yes = ok && slices.Contains(vDir, LEFT)
		}
	}
	if yes {
		// fmt.Printf("can block move at: %d, %d\n", g.r, g.c)
		g.loopCount++
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
	// fmt.Println(g.visited)
	fmt.Println("Part 2:", g.loopCount)
}
