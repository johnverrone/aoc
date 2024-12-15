package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
`

var re = regexp.MustCompile("p=([0-9]+),([0-9]+) v=(-?[0-9]+),(-?[0-9]+)")
var width = 101
var height = 103

type pos struct {
	x, y int
}

type robot struct {
	px, py int
	vx, vy int
}

func main() {
	in := util.ParseInput("")
	lines := strings.Split(in, "\n")

	robots := make([]*robot, 0, len(lines))

	for _, line := range lines {
		parts := re.FindAllStringSubmatch(line, -1)
		r := &robot{
			px: util.MustInt(parts[0][1]),
			py: util.MustInt(parts[0][2]),
			vx: util.MustInt(parts[0][3]),
			vy: util.MustInt(parts[0][4]),
		}
		robots = append(robots, r)
	}

	c := 1
	inARow := 0
	for {
		fmt.Println(c)

		// for range 100 {
		for _, r := range robots {
			r.move()
		}
		// }

		robotMap := map[pos]robot{}
		q1, q2, q3, q4 := 0, 0, 0, 0
		for _, r := range robots {
			hx := width / 2
			hy := height / 2
			if r.px < hx && r.py < hy {
				q1++
			}
			if r.px > hx && r.py < hy {
				q2++
			}
			if r.px < hx && r.py > hy {
				q3++
			}
			if r.px > hx && r.py > hy {
				q4++
			}
			robotMap[pos{r.px, r.py}] = *r
		}

		// sf := q1 * q2 * q3 * q4
		// fmt.Printf("%v\n", sf)

		// look for line
		for r := range height {
			for c := range width {
				if _, ok := robotMap[pos{c, r}]; ok {
					inARow++
				} else {
					inARow = 0
				}
				if inARow > 10 {
					print(robotMap)
					os.Exit(0)
				}
			}
		}
		c++
	}
}

func (r *robot) move() {
	nx, ny := r.px+r.vx, r.py+r.vy
	if nx < 0 {
		nx = width + nx
	}
	if nx >= width {
		nx = nx - width
	}
	if ny < 0 {
		ny = height + ny
	}
	if ny >= height {
		ny = ny - height
	}
	r.px = nx
	r.py = ny
}

func print(robots map[pos]robot) {
	for y := range height {
		for x := range width {
			if _, ok := robots[pos{x, y}]; ok {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
