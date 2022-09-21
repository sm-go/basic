package main

import "fmt"

type Counter struct {
	hits int
}

func myincrement(counter *Counter) {
	counter.hits += 1
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	myincrement(counter)
}

func main() {
	counter := Counter{}

	hello := "Hello"
	world := "World"
	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)
}
