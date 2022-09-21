package main

import (
	"fmt"
)

func test(x int) {
	fmt.Println("test function", x)
}

func add(a, b int) {
	fmt.Printf("Sum : %v + %v = ", a+b)
}

// return two numberes
func addret(x, y int) (int, int) {
	return x + y, x - y
}

func hiThere() string {
	return "Hi there!"
}

func main() {
	test(1)
	test(2)
	test(3)
	add(4, 5)
	fmt.Println(hiThere())

	ans1, ans2 := addret(12, 3)
	fmt.Println(ans1, ans2)
}
