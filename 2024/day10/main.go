package main

import (
	"fmt"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`

type pos struct {
	x, y int
}

type Trailhead struct {
	p     pos
	peaks []pos
}

type trailmap [][]int

func main() {
	in := util.ParseInput("")
	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")
	trail := make(trailmap, len(lines))

	// get trailheads
	var ths []Trailhead
	for y, row := range lines {
		tr := make([]int, len(row))
		for x, ch := range row {
			n := util.MustInt(string(ch))
			tr[x] = n
			if n == 0 {
				ths = append(ths, Trailhead{p: pos{x, y}})
			}
		}
		trail[y] = tr
	}

	sum := 0
	for _, th := range ths {
		trail.navigateTrail(&th, th.p)
		sum += len(th.peaks)
	}
	fmt.Printf("rating: %v\n", sum)
}

func (tm trailmap) navigateTrail(th *Trailhead, cur pos) {
	if tm[cur.y][cur.x] == 9 {
		// uncomment for pt 1
		// if !slices.Contains((*th).peaks, cur) {
		(*th).peaks = append((*th).peaks, cur)
		// }
	} else {
		next := tm.findNext(cur)
		for _, n := range next {
			tm.navigateTrail(th, n)
		}
	}
}

func (tm trailmap) findNext(p pos) []pos {
	var next []pos
	cur := tm[p.y][p.x]
	if p.y-1 >= 0 && tm[p.y-1][p.x] == cur+1 {
		// up is valid
		next = append(next, pos{p.x, p.y - 1})
	}
	if p.x+1 <= len(tm[0])-1 && tm[p.y][p.x+1] == cur+1 {
		// right is valid
		next = append(next, pos{p.x + 1, p.y})
	}
	if p.y+1 <= len(tm)-1 && tm[p.y+1][p.x] == cur+1 {
		// down is valid
		next = append(next, pos{p.x, p.y + 1})
	}
	if p.x-1 >= 0 && tm[p.y][p.x-1] == cur+1 {
		// left is valid
		next = append(next, pos{p.x - 1, p.y})
	}

	return next
}
