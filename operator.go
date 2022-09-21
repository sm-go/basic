package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 6.5
	val := float64(x) == y
	fmt.Printf("%t", val)
}
