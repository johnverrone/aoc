package main

import (
	"fmt"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `
AAAA
BBCD
BBCC
EEEC
`

const sample2 = `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
`

type pos struct {
	y, x int
}

type plot struct {
	p      pos
	plant  byte
	region int
}

func main() {
	in := util.ParseInput(sample2)
	lines := strings.Split(in, "\n")

	farm := map[int][]plot{}
	regionLookup := map[pos]int{}
	region := 0
	perimeterMap := map[int]int{}
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			cur := plot{
				p:     pos{y, x},
				plant: lines[y][x],
			}
			var plotAbove *plot
			var plotLeft *plot
			if y > 0 {
				plotAbove = &plot{
					p:      pos{y: y - 1, x: x},
					plant:  lines[y-1][x],
					region: regionLookup[pos{y: y - 1, x: x}],
				}
			}
			if x > 0 {
				plotLeft = &plot{
					p:      pos{y: y, x: x - 1},
					plant:  lines[y][x-1],
					region: regionLookup[pos{y: y, x: x - 1}],
				}
			}

			// check if plot matches existing regions in the farm
			if plotAbove != nil && plotAbove.plant == cur.plant && plotLeft != nil && plotLeft.plant == cur.plant {
				// found match and left
				cur.region = plotAbove.region
				fmt.Printf("found match above and left %d\n", perimeterMap[cur.region])
			} else if plotAbove != nil && plotAbove.plant == cur.plant {
				// found match above
				cur.region = plotAbove.region
				perimeterMap[cur.region] += 2
				fmt.Printf("found match above %d\n", perimeterMap[cur.region])
			} else if plotLeft != nil && plotLeft.plant == cur.plant {
				// found match left
				cur.region = plotLeft.region
				perimeterMap[cur.region] += 2
				fmt.Printf("found match left %d\n", perimeterMap[cur.region])
			} else {
				// no matches
				region++
				cur.region = region
				perimeterMap[cur.region] = 4
				fmt.Printf("no matches, increasing region id %d, setting perimeter %d\n", cur.region, perimeterMap[cur.region])
			}
			fmt.Printf("appending cur %v\n", cur)
			farm[cur.region] = append(farm[cur.region], cur)
			regionLookup[cur.p] = cur.region
		}
	}

	sum := 0
	for i, r := range farm {
		area := len(r)
		perimeter := perimeterMap[i]
		price := area * perimeter
		sum += price
		fmt.Printf("region %d (%c): %d * %d = %d\n", i, r[0].plant, area, perimeter, price)
	}
	fmt.Printf("%v\n", sum)
}
