package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 12; i++ {
		if i%5 == 0 {
			fmt.Printf("Hello World %d \n", i)
			break
		}
	}

	x := 1
	for x < 5 {
		fmt.Println(x)
		x++
	}
}
