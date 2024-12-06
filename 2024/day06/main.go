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

var lab []string

type Dir int

const (
	INVALID Dir = iota
	UP
	RIGHT
	DOWN
	LEFT
)

type point struct {
	r, c int
}

type Guard struct {
	r, c        int
	dir         Dir
	visited     map[point][]Dir
	blockPoints []point
}

func NewGuard(r, c int) *Guard {
	return &Guard{
		r:           r,
		c:           c,
		dir:         UP,
		visited:     map[point][]Dir{{r, c}: {UP}},
		blockPoints: []point{},
	}
}

func (g *Guard) move() {
	for {
		newPoint := point{g.r, g.c}
		for g.canMove() {
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
			// mark visited
			newPoint = point{g.r, g.c}
			fmt.Printf("moving to %v\n", newPoint)
			g.visited[newPoint] = append(g.visited[newPoint], g.dir)
		}
		// check bounds & turn
		if (g.dir == UP && g.r == 0) || (g.dir == DOWN && g.r == len(lab)-1) || (g.dir == RIGHT && g.c == len(lab[0])-1) || (g.dir == LEFT && g.c == 0) {
			break
		}
		if g.dir >= 4 {
			g.dir = 1
		} else {
			g.dir++
		}
		// add turn direction to visited
		g.visited[newPoint] = append(g.visited[newPoint], g.dir)
	}
}

func (g *Guard) canBlock() {
	yes := false
	switch g.dir {
	case UP:
		for i := g.c + 1; i <= len(lab[0])-1 && lab[g.r][i] != '#'; i++ {
			vDir, ok := g.visited[point{g.r, i}]
			yes = ok && slices.Contains(vDir, RIGHT)
		}
	case LEFT:
		for i := g.r - 1; i >= 0 && lab[i][g.c] != '#'; i-- {
			vDir, ok := g.visited[point{i, g.c}]
			yes = ok && slices.Contains(vDir, UP)
		}
	case RIGHT:
		for i := g.r + 1; i <= len(lab)-1 && lab[i][g.c] != '#'; i++ {
			vDir, ok := g.visited[point{i, g.c}]
			yes = ok && slices.Contains(vDir, DOWN)
		}
	case DOWN:
		for i := g.c - 1; i >= 0 && lab[g.r][i] != '#'; i-- {
			vDir, ok := g.visited[point{g.r, i}]
			yes = ok && slices.Contains(vDir, LEFT)
		}
	}
	if yes {
		blockPoint := getBlockPoint(g.dir, g)
		// fmt.Printf("can block move at: %v\n", blockPoint)
		// g.print(blockPoint)
		if !slices.Contains(g.blockPoints, blockPoint) {
			g.blockPoints = append(g.blockPoints, blockPoint)
		}
	}
}

func getBlockPoint(dir Dir, g *Guard) point {
	switch dir {
	case UP:
		return point{g.r - 1, g.c}
	case LEFT:
		return point{g.r, g.c - 1}
	case RIGHT:
		return point{g.r, g.c + 1}
	case DOWN:
		return point{g.r + 1, g.c}
	}
	panic("unhandled direction")
}

func (g *Guard) canMove() bool {
	switch g.dir {
	case UP:
		return g.r-1 >= 0 && lab[g.r-1][g.c] != '#'
	case LEFT:
		return g.c-1 >= 0 && lab[g.r][g.c-1] != '#'
	case RIGHT:
		return g.c+1 <= len(lab[0])-1 && lab[g.r][g.c+1] != '#'
	case DOWN:
		return g.r+1 <= len(lab)-1 && lab[g.r+1][g.c] != '#'
	}
	panic("unhandled move")
}

func main() {
	in := util.ParseInput("")
	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")
	lab = lines

	var g *Guard
	for r, row := range lines {
		c := strings.Index(row, "^")
		if c >= 0 {
			fmt.Printf("starting loc %d, %d\n", r, c)
			g = NewGuard(r, c)
		}
	}

	g.move()

	// part 1
	fmt.Printf("Part 1: %d\n", len(g.visited))

	// part 2
	fmt.Printf("Part 2: %d\n", len(g.blockPoints))
}

func (g *Guard) print(bp point) {
	for r, row := range lab {
		for c, col := range row {
			if c == bp.c && r == bp.r {
				fmt.Print("0")
				continue
			}
			if r == g.r && c == g.c {
				fmt.Print("^")
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
