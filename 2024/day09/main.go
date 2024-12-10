package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `2333133121414131402`

type chunk struct {
	start, size int
}

func (c chunk) String() string {
	return fmt.Sprintf("{ %d, %d }", c.start, c.size)
}

func main() {
	in := util.ParseInput("")
	in = strings.TrimSpace(in)

	blocks := []string{}
	emptySpot := []int{}
	emptyChunk := []*chunk{}
	file := true
	fileId := 0
	for _, ch := range in {
		size := util.MustInt(string(ch))
		if !file {
			emptyChunk = append(emptyChunk, &chunk{start: len(blocks), size: size})
		}
		for range size {
			if file {
				blocks = append(blocks, strconv.Itoa(fileId))
			} else {
				blocks = append(blocks, ".")
				emptySpot = append(emptySpot, len(blocks)-1)
			}
		}
		if file {
			fileId++
		}
		file = !file
	}
	pt2Blocks := make([]string, len(blocks))
	copy(pt2Blocks, blocks)

	// part 1
	times := len(emptySpot)
	for i := len(blocks) - 1; i >= len(blocks)-times; i-- {
		if blocks[i] == "." {
			continue
		}
		nextSpot := emptySpot[0]
		emptySpot = emptySpot[1:]
		blocks[nextSpot], blocks[i] = blocks[i], blocks[nextSpot]
	}

	sum := 0
	for i, b := range blocks {
		if b == "." {
			break
		}
		sum += util.MustInt(b) * i
	}
	fmt.Printf("Part 1: %v\n", sum)

	// part 2
	fileToMoveId := pt2Blocks[len(pt2Blocks)-1]
	fileToMoveSize := 0
	for i := len(pt2Blocks) - 1; i >= 0; i-- {
		if fileToMoveId == "." {
			fileToMoveId = pt2Blocks[i]
			continue
		} else if pt2Blocks[i] == fileToMoveId {
			fileToMoveSize++
		} else {
			// move it to open spot
			for _, chunk := range emptyChunk {
				if chunk.size >= fileToMoveSize && chunk.start < i+1 {
					emptyDots := make([]string, fileToMoveSize)
					for e := range fileToMoveSize {
						emptyDots[e] = "."
					}
					pt2Blocks = slices.Replace(pt2Blocks, chunk.start, chunk.start+fileToMoveSize, pt2Blocks[i+1:i+1+fileToMoveSize]...)
					pt2Blocks = slices.Replace(pt2Blocks, i+1, i+1+fileToMoveSize, emptyDots...)
					chunk.size -= fileToMoveSize
					chunk.start += fileToMoveSize
					break
				}
			}
			fileToMoveId = pt2Blocks[i]
			fileToMoveSize = 1
		}
	}

	sum = 0
	for i, b := range pt2Blocks {
		if b == "." {
			continue
		}
		sum += util.MustInt(b) * i
	}
	fmt.Printf("Part 2: %d\n", sum)
}
