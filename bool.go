package main

import (
	"fmt"
)

func main() {
	val := (true || false) && true
	fmt.Printf("%t", val)
}
