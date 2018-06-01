package main

import (
	"fmt"
	"math"
)

func sub(x, y int) int {
	return x - y
}

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("Now you have %g problems", math.Sqrt(6))
	fmt.Println("\n", add(12, 45))
	fmt.Println(sub(42, 12))
}
