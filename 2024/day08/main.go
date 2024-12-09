package main

import (
	"fmt"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`

type point struct {
	x, y     int
	antenna  *rune
	antinode bool
}

func (p point) String() string {
	return fmt.Sprintf("{ (%d, %d) antenna: '%s', antinode: %t }\n", p.x, p.y, runePtrToString(p.antenna), p.antinode)
}

func runePtrToString(p *rune) string {
	if p != nil {
		return string(*p)
	}
	return "(nil)"
}

func main() {
	in := util.ParseInput("")
	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")

	// part 1
	p, a := makePoints(lines, false)
	antinodeCount := 0
	// plot antinodes
	for _, v := range a {
		if len(v) < 2 {
			// no antinodes
			continue
		}

		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				// find distance between two points
				dx, dy := v[j].x-v[i].x, v[j].y-v[i].y

				x2 := v[j].x + dx
				y2 := v[j].y + dy
				if x2 >= 0 && x2 <= len(p[0])-1 && y2 >= 0 && y2 <= len(p)-1 {
					if !p[y2][x2].antinode {
						p[y2][x2].antinode = true
						antinodeCount++
					}
				}

				x3 := v[i].x - dx
				y3 := v[i].y - dy
				if x3 >= 0 && x3 <= len(p[0])-1 && y3 >= 0 && y3 <= len(p)-1 {
					if !p[y3][x3].antinode {
						p[y3][x3].antinode = true
						antinodeCount++
					}
				}
			}
		}
	}
	fmt.Printf("Part 1: %d\n", antinodeCount)

	// part 2
	p, a = makePoints(lines, true)
	antinodeCount = 0
	// plot antinodes
	for _, v := range a {
		if len(v) < 2 {
			// no antinodes
			continue
		}
		antinodeCount += len(v)

		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				// find distance between two points
				dx, dy := v[j].x-v[i].x, v[j].y-v[i].y

				x2 := v[j].x + dx
				y2 := v[j].y + dy
				for x2 >= 0 && x2 <= len(p[i])-1 && y2 >= 0 && y2 <= len(p)-1 {
					if !p[y2][x2].antinode {
						p[y2][x2].antinode = true
						antinodeCount++
					}
					x2 = x2 + dx
					y2 = y2 + dy
				}

				x3 := v[i].x - dx
				y3 := v[i].y - dy
				for x3 >= 0 && x3 <= len(p[i])-1 && y3 >= 0 && y3 <= len(p)-1 {
					if !p[y3][x3].antinode {
						p[y3][x3].antinode = true
						antinodeCount++
					}
					x3 = x3 - dx
					y3 = y3 - dy
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", antinodeCount)
}

func printPoints(points [][]*point) {
	for i := 0; i < len(points); i++ {
		fmt.Print(i, "\t")
		for j := 0; j < len(points[i]); j++ {
			if points[i][j].antinode {
				fmt.Print("#")
			} else if c := points[i][j].antenna; c != nil {
				fmt.Print(string(*c))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func makePoints(in []string, antennaIsAntinode bool) ([][]*point, map[rune][]*point) {
	points := make([][]*point, len(in))
	antennas := map[rune][]*point{}
	for r, line := range in {
		for c, char := range line {
			p := &point{
				x: c,
				y: r,
			}
			if char != '.' {
				p.antenna = &char
				if antennaIsAntinode {
					p.antinode = true
				}
			}
			if p.antenna != nil {
				antennas[*p.antenna] = append(antennas[*p.antenna], p)
			}
			points[r] = append(points[r], p)
		}
	}
	return points, antennas
}
