package main

import "fmt"

func runProgram1(testVars []int) {
	fmt.Println("=== Програма 1 (Одиничні порівняння) ===")
	for _, x := range testVars {
		y, ok := calcSimpleComparisons(float64(x))
		if ok {
			fmt.Println("x = %3d  -->  y = %8.3f\n", x, y)
		} else {
			fmt.Println("x = %3d  -->  функція не існує\n", x, y)
		}
	}
}

func runProgram2(testVars []int) {
	fmt.Println("=== Програма 2 (Булеві операції) ===")
	for _, x := range testVars {
		y, ok := calcBoolLogic(float64(x))
		if ok {
			fmt.Println("x = %3d  -->  y = %8.3f\n", x, y)
		} else {
			fmt.Println("x = %3d  -->  функція не існує\n", x, y)
		}
	}
}

func main() {
	testVars := []int{-5, -2, -1, 0, 4, 8, 9, 12, 15, 16, 25}
	runProgram1(testVars)
	runProgram2(testVars)
}
