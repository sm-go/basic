package main

import "fmt"

func increment(x *int) { // * is pointer
	*x += 1
}

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
	ii := 1
	myincrement(&ii) //& point to this i
	fmt.Println(ii)

	counter := Counter{}

	hello := "Hello"
	world := "World"
	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)
}
