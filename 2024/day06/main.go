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

var lab [][]byte

type Dir int

const (
	INVALID Dir = iota
	UP
	RIGHT
	DOWN
	LEFT
)

func (d Dir) String() string {
	switch d {
	case UP:
		return "UP"
	case RIGHT:
		return "RIGHT"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	}
	return "INVALID"
}

type point struct {
	r, c int
}

type Guard struct {
	start   point
	loc     point
	dir     Dir
	visited map[point][]Dir
}

func NewGuard(p point) *Guard {
	return &Guard{
		start:   p,
		loc:     p,
		dir:     UP,
		visited: map[point][]Dir{p: {UP}},
	}
}

type result int

const (
	DONE = result(iota)
	LOOP
)

func (g *Guard) move() result {
	// reset
	g.loc = g.start
	g.visited = map[point][]Dir{}
	g.dir = UP
	for {
		if g.isDone() {
			return DONE
		}
		if g.canMove() {
			if dir := g.visited[g.nextPoint()]; slices.Contains(dir, g.dir) {
				return LOOP
			}
			// move
			switch g.dir {
			case UP:
				g.loc.r--
			case LEFT:
				g.loc.c--
			case RIGHT:
				g.loc.c++
			case DOWN:
				g.loc.r++
			}
			// mark visited
			g.visited[g.loc] = append(g.visited[g.loc], g.dir)
		} else {
			// turn
			if g.dir >= 4 {
				g.dir = 1
			} else {
				g.dir++
			}
			// add turn direction to visited
			g.visited[g.loc] = append(g.visited[g.loc], g.dir)
		}
	}
}

func (g *Guard) nextPoint() point {
	switch g.dir {
	case UP:
		return point{g.loc.r - 1, g.loc.c}
	case LEFT:
		return point{g.loc.r, g.loc.c - 1}
	case RIGHT:
		return point{g.loc.r, g.loc.c + 1}
	case DOWN:
		return point{g.loc.r + 1, g.loc.c}
	}
	panic("unhandled direction")
}

func (g *Guard) isDone() bool {
	switch g.dir {
	case UP:
		return g.loc.r-1 < 0
	case LEFT:
		return g.loc.c-1 < 0
	case RIGHT:
		return g.loc.c+1 > len(lab[0])-1
	case DOWN:
		return g.loc.r+1 > len(lab)-1
	}
	panic("unhandled move")
}

func (g *Guard) canMove() bool {
	switch g.dir {
	case UP:
		return lab[g.loc.r-1][g.loc.c] != '#'
	case LEFT:
		return lab[g.loc.r][g.loc.c-1] != '#'
	case RIGHT:
		return lab[g.loc.r][g.loc.c+1] != '#'
	case DOWN:
		return lab[g.loc.r+1][g.loc.c] != '#'
	}
	panic("unhandled move")
}

func main() {
	in := util.ParseInput("")
	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")

	var start point
	lab = make([][]byte, len(lines))
	for r, line := range lines {
		for c, ch := range line {
			if ch == '^' {
				start = point{r, c}
			}
			lab[r] = append(lab[r], byte(ch))
		}
	}

	var g *Guard
	g = NewGuard(start)
	g.move()

	// part 1
	fmt.Printf("Part 1: %d\n", len(g.visited))

	// part 2
	sum := 0
	for r := 0; r < len(lab); r++ {
		for c := 0; c < len(lab[0]); c++ {
			p := point{r, c}
			if lab[r][c] == '.' && p != g.start {
				lab[r][c] = '#'

				res := g.move()
				if res == LOOP {
					sum++
				}

				lab[r][c] = '.'
			}
		}
	}
	fmt.Printf("Part 2: %d\n", sum)
}

func (g *Guard) print() {
	for r, row := range lab {
		for c, col := range row {
			if r == g.loc.r && c == g.loc.c {
				switch g.dir {
				case UP:
					fmt.Print("^")
				case DOWN:
					fmt.Print("v")
				case LEFT:
					fmt.Print("<")
				case RIGHT:
					fmt.Print(">")
				}
			} else if dir, ok := g.visited[point{r, c}]; ok {
				if len(dir) > 1 {
					fmt.Print("+")
				} else {
					switch dir[0] {
					case UP:
						fmt.Print("|")
					case DOWN:
						fmt.Print("|")
					case LEFT:
						fmt.Print("-")
					case RIGHT:
						fmt.Print("-")
					}
				}
			} else {
				fmt.Print(string(col))
			}
		}
		fmt.Println()
	}
}
