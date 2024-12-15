package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/johnverrone/aoc2024/util"
	"gonum.org/v1/gonum/mat"
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
		a1, b1, a2, b2 := util.MustFloat(moves[0][1]), util.MustFloat(moves[0][2]), util.MustFloat(moves[1][1]), util.MustFloat(moves[1][2])
		prize := pRe.FindAllStringSubmatch(game, -1)
		c1, c2 := util.MustFloat(prize[0][1]), util.MustFloat(prize[0][2])

		A := mat.NewDense(2, 2, []float64{a1, a2, b1, b2})
		b := mat.NewVecDense(2, []float64{c1, c2})

		var x mat.VecDense
		if err := x.SolveVec(A, b); err != nil {
			panic(err)
		}
		aTimes := -1
		bTimes := -1
		if x.AtVec(0) == math.Trunc(x.AtVec(0)) {
			aTimes = int(x.AtVec(0))
		}
		if x.AtVec(1) == math.Trunc(x.AtVec(1)) {
			bTimes = int(x.AtVec(1))
		}
		if aTimes > 100 || bTimes > 100 || aTimes < 0 || bTimes < 0 {
			continue
		}

		fmt.Printf("%d %d\n", aTimes*3, bTimes)
		sum += aTimes*3 + bTimes
	}

	fmt.Printf("%v\n", sum)
}
