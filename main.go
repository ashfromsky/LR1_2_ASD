package main

import (
	"fmt"
	"math"
)

const eps = 1e-9

func main() {
	testVars := []float64{-5, -1, -0.5, 0, 4, 8, 12, 16, 25}

	for _, x := range testVars {
		y1, ok1 := calcSimpleComparisons(x)
		y2, ok2 := calcBoolLogic(x)
		matched := (ok1 == ok2) && (!ok1 || math.Abs(y1-y2) < eps)

		if ok1 {
			fmt.Println(x, ":", y1, ":", y2, ":", matched)
		} else {
			fmt.Println(x, ": функція не існує :", matched)
		}
	}
}
