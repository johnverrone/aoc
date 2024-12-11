package main

import (
	"fmt"
	"maps"
	"strconv"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `125 17`
const input = `5910927 0 1 47 261223 94788 545 7771`

type stone int
type game struct {
	stones map[stone]int
}

func main() {
	in := util.ParseInput(input)
	in = strings.TrimSpace(in)
	ss := strings.Fields(in)

	// get stones
	g := &game{
		stones: make(map[stone]int, len(ss)),
	}
	for _, s := range ss {
		g.stones[stone(util.MustInt(s))] = 1
	}

	// 25 for pt 1
	for range 75 {
		g.tick()
	}

	sum := 0
	for s := range g.stones {
		sum += g.stones[s]
	}
	fmt.Printf("%v\n", sum)
}

func (g *game) tick() {
	oldStones := maps.Clone(g.stones)
	for s := range oldStones {
		if oldStones[s] <= 0 {
			continue
		}
		count := oldStones[s]
		g.stones[s] -= count
		tickStones := s.tick()
		for _, s := range tickStones {
			g.stones[s] += count
		}
	}
}

func (s stone) tick() []stone {
	if s == 0 {
		return []stone{1}
	}
	str := strconv.Itoa(int(s))
	if len(str)%2 == 0 {
		mid := len(str) / 2
		return []stone{stone(util.MustInt(str[:mid])), stone(util.MustInt(str[mid:]))}
	}
	return []stone{s * 2024}
}
