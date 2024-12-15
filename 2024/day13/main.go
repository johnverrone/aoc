package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279
`

var bRe = regexp.MustCompile(`X\+([0-9]*), Y\+([0-9]*)`)
var pRe = regexp.MustCompile(`X=([0-9]*), Y=([0-9]*)`)

func main() {
	in := util.ParseInput("")
	games := strings.Split(in, "\n\n")

	sum := 0
	for _, game := range games {
		moves := bRe.FindAllStringSubmatch(game, -1)
		ax, ay, bx, by := util.MustFloat(moves[0][1]), util.MustFloat(moves[0][2]), util.MustFloat(moves[1][1]), util.MustFloat(moves[1][2])
		prize := pRe.FindAllStringSubmatch(game, -1)
		x, y := util.MustFloat(prize[0][1]), util.MustFloat(prize[0][2])

		A := (bx*y - by*x) / (bx*ay - by*ax)
		B := (x - ax*A) / bx

		if A == math.Trunc(A) && B == math.Trunc(B) {
			a := int(A)
			b := int(B)
			sum += a*3 + b
		}
	}

	fmt.Printf("%v\n", sum)
}
